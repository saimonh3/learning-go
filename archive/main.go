package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// archive all the files and folders
	src := "."

	output, err := os.Create("file.zip")
	checkError(err)
	defer output.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		checkError(err)

		if info.IsDir() {
			return err
		}

		w := zip.NewWriter(output)
		defer w.Close()

		file, err := os.Open(path)
		checkError(err)

		defer file.Close()

		f, err := w.Create(path)
		checkError(err)

		_, err = io.Copy(f, file)
		checkError(err)

		return err
	})

	checkError(err)
}

func checkError(err error) {
	log.Fatal(err)
}