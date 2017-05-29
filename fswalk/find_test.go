package fswalk_test

import (
	"github.com/ghst659/grepgo/fswalk"
	"log"
	"testing"
)

func TestFlat(t *testing.T) {
	ch := make(chan string)
	go fswalk.Files("/tmp/foo", false, ch)
	i := 0
	for item := range ch {
		log.Printf("%d: %s", i, item)
		i++
	}
}

func TestDeep(t *testing.T) {
	ch := make(chan string)
	go fswalk.Files("/tmp/foo", true, ch)
	i := 0
	for item := range ch {
		log.Printf("%d: %s", i, item)
		i++
	}
}
