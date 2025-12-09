package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Bundler struct {
	indexPath string
	jsPaths   []string
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

func (bundler *Bundler) print() {
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

	bundler.print()

}
