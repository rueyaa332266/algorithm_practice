package main

import "fmt"

func reverse(slice []int) []int {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		slice[i], slice[j] = swap(slice[i], slice[j])
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
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(reverse(list))
}
