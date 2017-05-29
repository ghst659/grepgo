package fswalk

import (
	"os"
	"path/filepath"
)

func Files(root string, items chan<- string) {
	walker := func(path string, info os.FileInfo, err error) error {
		if ! info.IsDir() {
			items <- path
		}
		return nil
	}
	filepath.Walk(root, walker)
	close(items)
}
