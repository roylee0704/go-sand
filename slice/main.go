package main

import "fmt"

func main() {
	k := make([]string, 1)
	add()

	fmt.Println(k)

}

func add(t []string, v string) {
	t = append(t, v)
}
