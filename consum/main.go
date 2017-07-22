package main

import (
	"fmt"
)

func main() {
	list := []int{-1, 2, 4, -3, 5, 2, -5, 2}
	best := slowest(list)
	fmt.Println("best:", best)
}

// naive algo: O(n * n * n)
func slowest(list []int) int {
	n := len(list)
	best := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += list[k]
				best = max(sum, best)
			}
		}
	}
	return best
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
