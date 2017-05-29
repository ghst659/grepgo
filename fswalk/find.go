package fswalk

import (
	"os"
	"path/filepath"
	"strconv"
)

func Dummy(root string, recursive bool, items chan<- string) {
	for j := 15; j >= 0; j-- {
		items <- strconv.Itoa(j)
	}
	close(items)
}

func Files(root string, recursive bool, items chan<- string) {
	if recursive {
		walker := func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() { // regular file
				items <- path
			}
			return nil
		}
		filepath.Walk(root, walker)
	} else {
		items <- root
	}
	close(items)
}
