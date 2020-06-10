package main

import "fmt"

func insertion_sort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		slice = insertion(slice, i+1)
	}
	return slice
}

func insertion(slice []int, i int) []int {
	last := i - 1
	tmp := slice[last+1]
	for {
		if last < 0 {
			break
		} else if last >= 0 && slice[last] > tmp {
			slice[last+1] = slice[last]
			last = last - 1
		} else {
			break
		}
	}
	slice[last+1] = tmp
	return slice
}

func main() {
	list := []int{1, 9, 12, 5, 4, 10, 6, 8}
	fmt.Println(insertion_sort(list))
}
