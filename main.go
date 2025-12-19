package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Bundler struct {
	indexPath string
	jsPaths   []string
	jsContent string
}

func (bundler *Bundler) iterateFolder(dir string, indexName string) {

	var inputFolder, err = os.ReadDir(dir)

	if err != nil {
		log.Fatal("ERROR (read iterateFolder):", err)
	}

	for _, entry := range inputFolder {

		full := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			bundler.iterateFolder(full, indexName)
			continue
		}

		if entry.Name() == indexName && bundler.indexPath == "" {
			bundler.indexPath = full
			continue
		}

		if strings.HasSuffix(entry.Name(), ".js") {
			bundler.jsPaths = append(bundler.jsPaths, full)
		}

	}
}

/*
Current limitation:
  - Import's needs to be at the top of the file.
  - Only: Import ... from "..." \n
  - No require
*/
func (bundler *Bundler) lexer(keep bool) {

	if !keep {
		bundler.jsContent = ""
	}

	for _, entry := range bundler.jsPaths {
		var file, err = os.ReadFile(entry)
		if err != nil {
			log.Fatalf("Could not read %s: \n %s", entry, err)
		}

		var tempBuf bytes.Buffer

		tempBuf.Write(file)

		var fileContent = tempBuf.String()

		for _, line := range strings.Split(fileContent, "\n") {
			trimmed := strings.TrimSpace(line)
			if len(trimmed) > 0 {
				if strings.HasPrefix(trimmed, "import") {
					continue
				}

				if strings.HasPrefix(trimmed, "export") {
					trimmed = trimmed[7:]
				}

				bundler.jsContent += trimmed
				bundler.jsContent += "\n"
			}
		}
	}

}

func (bundler *Bundler) printContent() {
	fmt.Printf("%s \n", bundler.jsContent)
}

func (bundler *Bundler) printPaths() {
	fmt.Printf("Index HTML File: \n")
	fmt.Printf("    %s \n", bundler.indexPath)
	fmt.Printf("JS Files: \n")
	for idx, entry := range bundler.jsPaths {
		fmt.Printf("    %d:  %s\n", idx, entry)
	}
}

func main() {

	const inputSrc = "example"
	const outputSrc = "output"

	const indexName = "index.html"

	var bundler = Bundler{
		indexPath: "",
	}

	bundler.iterateFolder(inputSrc, indexName)

	bundler.lexer(false)

	bundler.printContent()
	// bundler.printPaths()

}
