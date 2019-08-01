package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h      bool
	config string
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: mqttagent [-h] [-c filename]

Options:
`)
	flag.PrintDefaults()
}
func init() {
	flag.BoolVar(&h, "h", false, "get help")
	flag.StringVar(&config, "c", "./mqttagent.ini", "set configuration `file`")
	flag.Usage = usage
}
func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	//your logic
}
