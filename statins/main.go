package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	for _, item := range flag.Args() {
		info, err := os.Stat(item)
		fmt.Println(item)
		fmt.Println(err)
		fmt.Println(info.Mode())
		fmt.Println(info.Mode().IsRegular())
		fmt.Println("###############")
	}
}
