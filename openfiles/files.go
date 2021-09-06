package openfiles

import (
	"fmt"
	"os"

	"github.com/Tum/logga/writelogs"
	"github.com/mitchellh/go-homedir"
)

func FindHomeDir() string {
	home, err := homedir.Dir()
	check(err)
	return home
}

func DirLocationAndName() string {
	name := FindHomeDir()
	dirName := fmt.Sprintf("%s/%s", name, ".glearn")
	return dirName
}

func ReadHomeDir(home string) {
	dir, err := os.ReadDir(home)
	check(err)
	dirName := DirLocationAndName()

	for _, fl := range dir {
		if fl.Name() == ".glearn" {
			return
		}

		err = os.Mkdir(dirName, 0755)
		check(err)
	}
}

func check(err error) {
	lg := new(writelogs.WriteLogs)
	if err != nil {
		lg.WriteIntoLogFiles(err)
	}
}
