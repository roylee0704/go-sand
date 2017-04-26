package main

import (
	"fmt"
	"time"
)

func main() {
	x := 100
	fmt.Printf("calculating fib(%d)...\n", x)
	go spinner(100 * time.Millisecond)
	fmt.Println(fib(x))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
