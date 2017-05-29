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
			if hits, err := grep.SearchFile(target); err != nil {
				log.Println(err)
			} else {
				for _, hit := range hits {
					fmt.Printf("%s:%d:%s", target, hit.Line, hit.Text)
				}
			}
		}
	} else {
		log.Println(err)
	}
}
