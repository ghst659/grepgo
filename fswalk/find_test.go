package fswalk_test

import "fmt"
import "testing"
import "github.com/ghst659/grepgo/fswalk"

func TestSequential(t *testing.T) {
	ch := make(chan string)
	go fswalk.Files("/tmp/foo", fswalk.Series, ch)
	for i := range ch {
		fmt.Println(i)
	}
}
