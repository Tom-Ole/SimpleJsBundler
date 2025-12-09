package main

import (
	"fmt"
	"log"
	"os"
)

type Bundler struct {
	indexFile *os.File
	jsFiles   []*os.File
}

func (bundler *Bundler) iterateFolder(inputSrc string, indexName string, pathAcc string) {

	var path = pathAcc + inputSrc

	var inputFolder, err = os.ReadDir(path)

	if err != nil {
		log.Fatal("ERROR (19):", err)

	}

	for _, entry := range inputFolder {
		if entry.IsDir() {
			bundler.iterateFolder(entry.Name(), indexName, path+"/")
		}

		if entry.Name() == indexName {
			if bundler.indexFile != nil {
				continue
			}

			var file, err = os.Open(inputSrc + "/" + entry.Name())
			defer file.Close()

			if err != nil {
				log.Fatal("ERROR (36):", err)
			}

			bundler.indexFile = file
		}

		if entry.Name()[(len(entry.Name())-3):] == ".js" {
			var file, err = os.Open(path + "/" + entry.Name())
			defer file.Close()

			if err != nil {
				log.Fatal("ERROR (47):", err)
			}

			bundler.jsFiles = append(bundler.jsFiles, file)
		}

	}
}

func (bundler *Bundler) print() {
	fmt.Printf("Index HTML File: \n")
	fmt.Printf("    %s \n", bundler.indexFile.Name())
	fmt.Printf("JS Files: \n")
	for idx, entry := range bundler.jsFiles {
		fmt.Printf("    %d:  %s\n", idx, entry.Name())
	}
}

func main() {

	const inputSrc = "example"
	const outputSrc = "output"

	const indexName = "index.html"

	var bundler = Bundler{
		indexFile: nil,
	}

	bundler.iterateFolder(inputSrc, indexName, "")

	bundler.print()

}
