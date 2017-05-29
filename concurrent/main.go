package main

import (
	"fmt"
	"github.com/ghst659/grepgo/grepper"
	"log"
)

func main() {
	grep, targets, err := grepper.CliParse()
	if err == nil {
		for target := range targets {
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
