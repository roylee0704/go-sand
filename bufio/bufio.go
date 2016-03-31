package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	reader := strings.NewReader("abdef")
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
