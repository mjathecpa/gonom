package gonomfunc

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	s "strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
)

const yyyymmRegx := 

func GetFiles() []string {
	var files []string

	root := SelectFolder()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func SelectFolder() string {
	pickFolderDialog, err := cfd.NewSelectFolderDialog(cfd.DialogConfig{
		Title: "Pick Folder",
		Role:  "PickFolderExample",
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := pickFolderDialog.Show(); err != nil {
		log.Fatal(err)
	}
	result, err := pickFolderDialog.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Chosen folder: %s\n", result)

	return result
}

func Mv(oldName string, newName string) {
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func MvExt(origFile string, oldExt string, newExt string) {
	// check if file meets suffix text
	if s.HasSuffix(origFile, oldExt) {
		// replace original filename with replaced extenstion filename
		err := os.Rename(origFile, s.Replace(origFile, oldExt, newExt, -1))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func MvRxp(origFile string, oldVal string, newVal string){
	re := regexp.MustCompile( "([2]\d\d\d)(\-)([0-1]\d)\s" )
	newName := re.ReplaceAllString(oldVal, newVal)
}
