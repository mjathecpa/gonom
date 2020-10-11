package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	sample1 := "2020-09 Changer 2002-087.pdf"
	//sample2 := "Don't change 2020-09 -087.pdf"
	newStrg := "2020-10"
	// compile search parameter
	reParam, err := regexp.Compile(`^([2]\d\d\d)(\-)([0-1]\d)`)
	if err != nil {
		log.Fatal(err)
	}
	// identify if rexp

	found := reParam.MatchString(sample1)
	fmt.Println("Sample1 found? %s", strconv.FormatBool(found))

	if found {
		newName := reParam.ReplaceAllString(sample1, newStrg)
		// replace original filename with replaced extenstion filename
		fmt.Println("Changed Sample1 = %s", newName)
		// errr := os.Rename(origFile, newName)
		// if errr != nil {
		// 	log.Fatal(errr)
		// }
	}
}
