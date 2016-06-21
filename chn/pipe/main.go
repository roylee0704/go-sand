package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarers := make(chan int)

	go natural(naturals)
	go squarer(naturals, squarers)
	printer(squarers)
}

// natural number generator
func natural(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out) //asserts no more sends.
}

// square number generator
func squarer(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

// printer prints the output
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
