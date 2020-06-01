package main

import "fmt"

func minValue(in []int) int {
	// int64 max number
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
	fmt.Println(minValue(list))
}
