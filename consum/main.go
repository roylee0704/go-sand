package main

import (
	"fmt"
)

func main() {
	list := []int{-1, 2, 4, -3, 5, 2, -5, 2}
	best := slowest(list)
	best = medium(list)
	best = efficient(list)

	fmt.Println("best:", best)
}

// o(n). largest sum: sequential-sum + array[point] > array[point]
func efficient(list []int) int {
	n := len(list)
	best := 0
	sum := 0
	var collection []int
	for i := 0; i < n; i++ {
		sum = max(list[i], sum+list[i])
		if sum == list[i] {
			collection = nil
		}
		collection = append(collection, list[i])
		best = max(sum, best)
	}
	fmt.Println(collection)
	return best
}

//  o(n * n)
// 1,2,3,4,5
// 2,3,4,5
// 3,4,5,
// 4,5
// 5,
func medium(list []int) int {
	n := len(list)
	best := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += list[j]
			best = max(sum, best)
		}
	}
	return best
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
