package main

import "fmt"

func merge(slice []int, l int, m int, r int) []int {
	var mergeSlice []int
	tmp := append(slice[l:m], reverse(slice[m:r])...)
	left := 0
	right := r - l - 1

	for i := l; i <= r-1; i++ {
		if tmp[left] <= tmp[right] {
			mergeSlice = append(mergeSlice, tmp[left])
			left = left + 1
		} else {
			mergeSlice = append(mergeSlice, tmp[right])
			right = right - 1
		}
	}

	for i := range slice[l:r] {
		slice[l:r][i] = mergeSlice[i]
	}
	return slice
}

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
	list := []int{17, 22, 39, 44, 53, 68, 76, 2, 45, 53, 64, 79, 88}
	fmt.Println(merge(list, 1, 6, 13))
}
