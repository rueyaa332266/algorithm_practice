package main

import "fmt"

// return target index or -1(not found)
func binarySearch(in []int, target int) int {
	left := 0
	right := len(in) - 1
	for left < right {
		var mid int
		mid = (right + left) / 2
		// Check left and right if there is no value between them
		if mid < 1 {
			if in[left] == target {
				return left
			} else if in[right] == target {
				return right
			} else {
				return -1
			}
		}

		if in[mid] == target {
			return mid
		} else if in[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarySearch(list, 6))
}
