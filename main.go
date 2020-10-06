package main

import (
	"fmt"
	f "gonom/gonomfunc"
	"path/filepath"
)

func main() {
	files := f.GetFiles()

	for _, file := range files {
		fileName := filepath.Base(file)
		fmt.Println(fileName)
	}
}
