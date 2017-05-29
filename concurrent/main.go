package main

import (
	"fmt"
	"log"
	"github.com/ghst659/grepgo/grepper"
)

func main() {
	grep, targets, _, err := grepper.CliParse()
	if err == nil {
		for _, target := range targets {
			results := make(chan grepper.SearchHit)
			go grep.YieldHits(target, results)
			for hit, not_done := <-results; not_done; hit, not_done = <-results {
				fmt.Printf("%s:%d:%s", target, hit.Line, hit.Text)
			}
		}
	} else {
		log.Println(err)
	}
}
