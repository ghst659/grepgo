package fswalk

import (
	"log"
	"os"
	"path/filepath"
)

type WalkMode int

const (
	NoWalk WalkMode = iota
	Series
	Parallel
)

func Files(root string, mode WalkMode, items chan<- string) {
	walker = func(path string, info os.FileInfo, err error) error {
		currentPath = filepath.Join(root, info.Name())
		if info.IsDir() {
			switch mode {
			case Series:
				Files(currentPath, mode, items)
			case Parallel:
				go Files(currentPath, mode, items)
			default:
				log.Printf("%s: no action", currentPath)
			}
		} else {
			items <- currentPath
		}
		return nil
	}
	filepath.Walk(root, walker)
	close(items)
}
