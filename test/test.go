package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	filepath.Walk(dir, func(path string, file fs.FileInfo, err error) error {
		fmt.Println(file.Name())
		return nil
	})
}
