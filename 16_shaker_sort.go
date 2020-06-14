package main

import "fmt"

func shakerSort(slice []int) []int {
	left := 0
	right := len(slice) - 1
	for {
		checkIndex := 0
		// left to right
		for i := left + 1; i <= right; i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = swap(slice[i-1], slice[i])
				checkIndex = i - 1
			}
		}
		right = checkIndex

		if left == right {
			break
		}
		// right to left
		for j := right; j > left; j-- {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = swap(slice[j], slice[j-1])
				checkIndex = j
			}
		}
		left = checkIndex
		if left == right {
			break
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
	list := []int{3, 1, 2, 12, 10, 5, 7, 4, 9, 11}
	fmt.Println(shakerSort(list))
}
