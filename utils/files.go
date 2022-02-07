package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func CheckFileAlreadyExists(file string) bool {
	fs, err := os.Lstat(file)

	if err != nil {
		return false
	}

	if fs != nil {
		return true
	}

	return false
}

func CheckDirExists(file string) bool {
	dir := filepath.Dir(file)
	_, err := os.ReadDir(dir)

	switch e := err.(type) {
	default:
		fmt.Printf("An unspecified error occured: %s", e.Error())
		os.Exit(3)
		return false

	case *fs.PathError:
		return false

	case nil:
		return true
	}

}
