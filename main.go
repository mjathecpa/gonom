package main

import (
	"fmt"
	f "gonom/gonomfunc"
	s "strings"
)

func main() {

	var userOpt string
	var findVal string
	var replVal string

	fmt.Println(userOpt)

	for {
		fmt.Println("Select a program option:")
		fmt.Println("---------------------")
		fmt.Println("1 - Change file extensions")
		fmt.Println("2 - Change date reference (yyyy-mm)")
		fmt.Scan(&userOpt)

		if userOpt == "1" {
			// get find/replace values
			fmt.Println("Enter extension to find:")
			fmt.Scan(&findVal)
			fmt.Println("Enter extension to replace:")
			fmt.Scan(&replVal)
			// set extension to lowercase
			replVal = s.ToLower(replVal)

			// query user for file name array
			files := f.GetFiles()

			// cycle and replace
			for _, file := range files {
				f.MvExt(file, findVal, replVal)
			}
		} else if userOpt == "2" {
			// get find/replace values
			fmt.Println("Enter date prefix to replace:")
			fmt.Scan(&replVal)

			// query user for file name array
			files := f.GetFiles()

			// cycle and replace
			for _, file := range files {
				f.MvYyyyMmRxp(file, replVal)
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
