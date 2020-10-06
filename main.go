package main

import (
	"fmt"
	f "gonom/gonomfunc"
	s "strings"
)

func main() {

	var userOpt string

	fmt.Println(userOpt)

	for {
		fmt.Println("Select a program option:")
		fmt.Println("---------------------")
		fmt.Println("1 - Change file extensions")
		fmt.Scan(&userOpt)

		if userOpt == "1" {
			// get find/replace values
			var findExt string
			var replExt string
			fmt.Println("Enter extension to find:")
			fmt.Scan(&findExt)
			fmt.Println("Enter extension to replace:")
			fmt.Scan(&replExt)
			replExt = s.ToLower(replExt)

			// get file name array
			files := f.GetFiles()

			// cycle and replace
			for _, file := range files {
				f.MvExt(file, findExt, replExt)
			}

		} else {
			fmt.Println("No valid option selected")
		}
		fmt.Println("Process another (y/n)")
		fmt.Scan(&userOpt)
		if s.ToLower(userOpt) == "n" {
			break
		}
	}

	/*
		files := f.GetFiles()
		for _, file := range files {
			fileName := filepath.Base(file)
			fmt.Println(fileName)
		}*/
}
