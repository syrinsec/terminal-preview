package main

import (
	"flag"
	"fmt"
  "io/ioutil"
)

// Variables used for command line parameters
var (
	Path string
)
func init() {
	flag.StringVar(&Path, "f", "", "File Path")
	flag.Parse()
}

func main() {
  data, err := ioutil.ReadFile(Path)
  if err != nil {
    fmt.Println("Couldn\'t find file", err)
    return
  }
  fmt.Println(string(data))
}
