package main

import "fmt"

func bubble_sort(slice []int) []int {
	left := 0
	right := len(slice) - 1
	for i := left; i < right; i++ {
		for j := right; j > i; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = swap(slice[j-1], slice[j])
			}
		}
	}
	return slice
}

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	list := []int{1, 9, 12, 5, 4, 10, 6, 8}
	fmt.Println(bubble_sort(list))
}
