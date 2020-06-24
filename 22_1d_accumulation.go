package main

import (
	"fmt"
)

type segment struct {
	left  int
	right int
}

func accumulation1d(in []int, set []segment) []int {
	for _, v := range set {
		in[v.left] = in[v.left] + 1
		in[v.right] = in[v.right] - 1
	}
	for i := 1; i <= (len(in) - 1); i++ {
		in[i] = in[i] + in[i-1]
	}
	return in

}

func main() {
	seg1 := segment{1, 2}
	seg2 := segment{4, 7}
	seg3 := segment{2, 6}
	seg4 := segment{3, 5}
	set := []segment{seg1, seg2, seg3, seg4}
	list := make([]int, 8)
	fmt.Println(accumulation1d(list, set))
}
