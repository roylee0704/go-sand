package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	fmt.Println("Rocket launched!")

}
