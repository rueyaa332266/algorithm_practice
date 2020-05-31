package main

import "fmt"

func min(in []int) int {
	min := 9223372036854775807
	for i := range in {
		if min > in[i] {
			min = in[i]
		}
	}
	return min
}

func main() {
	list := []int{2, 3, 4, 3, 2, 20}
	fmt.Println(min(list))
}
