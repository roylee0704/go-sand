package main

import "fmt"

func main() {

	s := "สถานีรถไ"

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	s2 := "สถานี���ถไ"

	for i, r := range s2 {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
