package main

import "fmt"

func partition(slice []int, start int, end int) []int {
	s := start - 1
	for i := start; i < end; i++ {
		if slice[i] < slice[end] {
			s = s + 1
			slice[s], slice[i] = swap(slice[s], slice[i])
		}
	}
	s = s + 1
	slice[s], slice[end] = swap(slice[s], slice[end])

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
	fmt.Println(partition(list, 0, len(list)-1))
}
