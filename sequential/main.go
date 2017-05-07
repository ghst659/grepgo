package main

import (
  "fmt"
  "flag"
  "log"
  "github.com/ghst659/grepgo/grepper"
)

func main() {
  var fp_fixed = flag.Bool("fixed", false, "fixed expression")
  flag.Parse()
  var grep grepper.Grepper
  var targets []string
  for i, arg := range flag.Args() {
    if i == 0 {
      var err error
      if grep, err = grepper.NewGrepper(arg, *fp_fixed); err != nil {
        log.Println(err)
      }
    } else {
      targets = append(targets, arg)
    }
  }
  //***********************
  for _, target := range targets {
    if hits, err := grep.SearchFile(target); err != nil {
      log.Println(err)
    } else {
      for _, hit := range hits {
        fmt.Printf("%s:%d:%s", target, hit.Line, hit.Text)
      }
    }
  }
}
