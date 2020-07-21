package main

import "fmt"

func countingSort(list []int) []int {
	count := make([]int, len(list)-1)
	out := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		count[list[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}

	for i := len(list) - 1; i >= 0; i-- {
		count[list[i]]--
		out[count[list[i]]] = list[i]
	}

	return out
}

func main() {
	list := []int{3, 2, 4, 3, 4, 1}
	fmt.Println(countingSort(list))
}
