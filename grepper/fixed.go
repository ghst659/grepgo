package grepper

import (
	"strings"
)

type fgrep struct {
	expression string
}

func (fgrep *fgrep) SearchFile(path string) ([]SearchHit, error) {
	checker := func(text string) bool {
		return strings.Contains(text, fgrep.expression)
	}
	return searchLoop(path, checker)
}

func (fgrep *fgrep) YieldHits(path string, results chan<- SearchHit) {
	checker := func(text string) bool {
		return strings.Contains(text, fgrep.expression)
	}
	yieldLoop(path, checker, results)
}
