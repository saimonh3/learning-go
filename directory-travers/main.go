package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {

	root := "/Users/saimon/Documents/www/blog"
	timer := time.Now()
	err := traversByWalk(root)
	checkError(err)
	err = traversByReadDir(root)
	checkError(err)
	end := time.Now()

	fmt.Println(end.Sub(timer))
}

func traversByWalk(path string) error {
	return filepath.Walk(path, visit)
}

func visit(path string, file os.FileInfo, err error) error {
	checkError(err)
	fmt.Println(path, file.Size())

	return err
}

func traversByReadDir(path string) error {
	files, err := ioutil.ReadDir(path)

	checkError(err)

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return err
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

