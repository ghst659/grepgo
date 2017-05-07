package grepper

import (
	"regexp"
)

type egrep struct {
	pattern *regexp.Regexp
}

func (egrep *egrep) SearchFile(path string) ([]SearchHit, error) {
	checker := func(text string) bool {
		return egrep.pattern.FindString(text) != ""
	}
	return searchLoop(path, checker)
}

func (egrep *egrep) YieldHits(path string, results chan<- SearchHit) {
	checker := func(text string) bool {
		return egrep.pattern.FindString(text) != ""
	}
	yieldLoop(path, checker, results)
}
