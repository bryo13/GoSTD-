package main

import (
	"errors"

	"github.com/Tum/logga/openfiles"
	"github.com/Tum/logga/writelogs"
)

func main() {

	lg := new(writelogs.WriteLogs)
	home := openfiles.FindHomeDir()
	openfiles.ReadHomeDir(home)
	lg.WriteIntoLogFiles(errors.New("Final season"))
}
