package main

import (
	"fmt"
)

func euclideanSlgorithm(a int, b int) int {
	for b > 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}

func main() {
	fmt.Println(euclideanSlgorithm(36, 16))
}
