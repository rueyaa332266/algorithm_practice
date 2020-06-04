package main

import "fmt"

// Insert is the last element
func insertion(slice []int) []int {
	last := len(slice) - 2
	tmp := slice[len(slice)-1]
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
	list := []int{10, 50, 20, 60, 100, 8}
	fmt.Println(insertion(list))
}
