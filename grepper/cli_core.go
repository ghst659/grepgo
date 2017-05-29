package grepper

import (
	"flag"
	"github.com/ghst659/grepgo/fswalk"
	"log"
	"os"
)

func CliParse() (grep Grepper, targets chan string, err error) {
	fpRecursive := flag.Bool("recursive", false, "process whole directory trees")
	fpFixed := flag.Bool("fixed", false, "fixed expression")
	flag.Parse()

	pattern := flag.Arg(0)
	grep, err = NewGrepper(pattern, *fpFixed)

	if err == nil {
		targets = make(chan string)
		go getTargets(flag.Args()[1:], targets, *fpRecursive)
	}
	return
}

func getTargets(items []string, targets chan string, recursive bool) {
	for _, item := range items {
		_, code := os.Stat(item)
		if os.IsNotExist(code) {
			log.Printf("%s: no such file or directory", item)
		} else {
			fswalk.Files(item, recursive, targets)
			// targets <- item
		}
	}
	// close(targets)
}
