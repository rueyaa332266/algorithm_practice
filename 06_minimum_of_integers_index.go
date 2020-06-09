package main

import "fmt"

func minIndex(in []int) int {
	mini := 0
	for i := range in {
		if in[i] < in[mini] {
			mini = i
		}
	}
	return mini
}

func main() {
	list := []int{244, 3, 4, 1, 3, 2, 20}
	fmt.Println(minIndex(list))
}
