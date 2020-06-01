package main

import "fmt"

// return match index or -1(not found)
func linear_search(in []int, target int) int {
	for i := range in {
		if in[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	list := []int{244, 3, 4, 1, 3, 2, 20}
	fmt.Println(linear_search(list, 3))
}
