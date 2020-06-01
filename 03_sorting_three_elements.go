package main

import "fmt"

func sort3(in []int) []int {
	for i := range in {
		if i+1 == len(in) {
			break
		}
		if in[i] > in[i+1] {
			in[i], in[i+1] = swap(in[i], in[i+1])
		}
	}
	return in
}

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	test := []int{2, 5, 3, 100, 4, 55}
	fmt.Println(sort3(test))
}
