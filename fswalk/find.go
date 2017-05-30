package fswalk

import (
	"log"
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

func Files(item string, recursive bool, ch chan<- string) {
	if recursive {
		walker := func(path string, info os.FileInfo, err error) error {
			if info.Mode().IsRegular() {
				ch <- path
			}
			return nil
		}
		filepath.Walk(item, walker)
	} else {
		info, err := os.Stat(item)
		if err == nil && info.Mode().IsRegular() {
			ch <- item
		} else {
			log.Printf("%s: not a file\n", item)
		}
	}
}
