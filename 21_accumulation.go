package main

import (
	"fmt"
)

func accumulation(in []int, left int, right int) int {
	ac := make([]int, len(in)+1)
	ac[0] = 0
	for i, v := range in {
		ac[i+1] = ac[i] + v
	}
	return ac[right] - ac[left-1]

}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(accumulation(list, 1, 4))
}
