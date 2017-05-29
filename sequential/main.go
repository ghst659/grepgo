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
