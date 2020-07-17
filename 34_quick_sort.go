package main

import "fmt"

func quickSort(slice []int, l int, r int) []int {
	if l < r {
		_, q := partition(slice, l, r)
		fmt.Println(q)
		quickSort(slice, l, q-1)
		quickSort(slice, q+1, r)
	}
	return slice
}

func partition(slice []int, start int, end int) ([]int, int) {
	s := start - 1
	for i := start; i < end; i++ {
		if slice[i] < slice[end] {
			s = s + 1
			slice[s], slice[i] = swap(slice[s], slice[i])
		}
	}
	s = s + 1
	slice[s], slice[end] = swap(slice[s], slice[end])

	return slice, s
}

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	list := []int{21, 14, 11, 2, 5, 18, 1, 25, 9, 6, 3, 7}
	fmt.Println(quickSort(list, 0, 11))
}
