package main

import "fmt"

func merge(slice1 []int, slice2 []int) []int {
	var mergeSlice []int
	tmp := append(slice1, reverse(slice2)...)
	left := 0
	right := len(tmp) - 1

	for {
		if left == right {
			mergeSlice = append(mergeSlice, tmp[left])
			break
		} else if tmp[left] < tmp[right] {
			mergeSlice = append(mergeSlice, tmp[left])
			left = left + 1
		} else {
			mergeSlice = append(mergeSlice, tmp[right])
			right = right - 1
		}
	}

	return mergeSlice
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
	list1 := []int{17, 22, 39, 44, 53, 68, 76}
	list2 := []int{2, 45, 53, 64, 79, 88}
	fmt.Println(merge(list1, list2))
}
