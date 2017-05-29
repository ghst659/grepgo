package fswalk_test

import (
	"log"
	"testing"
	"github.com/ghst659/grepgo/fswalk"
)

func TestSequential(t *testing.T) {
	ch := make(chan string)
	go fswalk.Files("/tmp/foo", ch)
	i := 0
	for item := range ch {
		log.Printf("%d: %s", i, item)
		i++
	}
}
