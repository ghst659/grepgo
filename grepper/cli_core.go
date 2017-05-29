package grepper

import (
	"flag"
)

func CliParse() (grep Grepper, targets []string, recursive bool, err error) {
	fp_recursive := flag.Bool("recursive", false, "process whole directory trees")
	fp_fixed := flag.Bool("fixed", false, "fixed expression")
	flag.Parse()
	recursive = *fp_recursive
	pattern := flag.Arg(0)
	grep, err = NewGrepper(pattern, *fp_fixed)
	targets = flag.Args()[1:]
	return
}
