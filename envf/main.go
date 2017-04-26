package main

import (
	"flag"
	"fmt"

	"github.com/gobike/envflag"
)

func main() {
	var times int

	flag.IntVar(&times, "f-times", 1, "this is a number")
	envflag.Parse()
	fmt.Println(times)
}
