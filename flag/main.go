package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "seperator")
var n = flag.Bool("n", false, "omit trailing newline")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
