package grepper

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

//******************************************************************************
type SearchHit struct {
	Line uint
	Text string
}
//******************************************************************************
type Grepper interface {
	SearchFile(path string) ([]SearchHit, error)
	YieldHits(path string, results chan<- SearchHit)
}

func NewGrepper(spec string, fixed bool) (result Grepper, err error) {
	var pattern *regexp.Regexp
	if fixed {
		result = &fgrep{spec}
	} else if pattern, err = regexp.Compile(spec); err == nil {
		result = &egrep{pattern}
	} else {
		// do nothing; err is already set
	}
	return
}
//******************************************************************************
func yieldLoop(path string, filter func(string) bool, results chan<- SearchHit) {
	defer close(results)
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		source := bufio.NewReader(file)
		for line_num := uint(0); err == nil; line_num++ {
			if text, err := source.ReadString('\n'); err == nil && filter(text) {
				results <- SearchHit{line_num, text}
			} else if err == io.EOF {
				break
			}
		}
	}
}

func searchLoop(path string, filter func(string) bool) (result []SearchHit, err error) {
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		source := bufio.NewReader(file)
		for line_num := uint(0); err == nil; line_num++ {
			if text, err := source.ReadString('\n'); err == nil && filter(text) {
				result = append(result, SearchHit{line_num, text})
			} else if err == io.EOF {
				err = nil
				break
			} else {
				// do nothing
			}
		}
	}
	return
}
