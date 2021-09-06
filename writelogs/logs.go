package writelogs

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
)

type WriteLogs struct {
	L    *log.Logger
	File *os.File
}

func dir() string {
	home, _ := homedir.Dir()
	glearn := fmt.Sprintf("%s/%s", home, ".glearn")
	return glearn
}

func (l *WriteLogs) WriteIntoLogFiles(err error) *log.Logger {
	path := dir()
	filepathandname := fmt.Sprintf("%s/%s", path, " logga.log")
	l.File, _ = os.OpenFile(filepathandname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	l.L = log.New(l.File, "logga", log.Llongfile)
	log.SetOutput(l.File)
	l.L.Println(err)

	defer l.File.Close()

	return l.L

}
