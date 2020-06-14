package main

import "fmt"

func selectionSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		index := minIndex(slice[i+1:])
		if slice[i] > slice[i+index+1] {
			slice[i], slice[i+index+1] = swap(slice[i], slice[i+index+1])
		}
	}
	return slice
}

func minIndex(in []int) int {
	mini := 0
	for i := range in {
		if in[i] < in[mini] {
			mini = i
		}
	}
	return mini
}

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	list := []int{1, 9, 12, 5, 4, 10, 6, 8}
	fmt.Println(selectionSort(list))
}
