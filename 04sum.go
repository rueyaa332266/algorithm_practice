package main

import "fmt"

func sum(in []int) int {
	var sum int
	for i := range in {
		sum += in[i]
	}
	return sum
}

func main() {
	list := []int{1, 2, 3, 4, 3, 2, 1, 20}
	fmt.Println(sum(list))
}
