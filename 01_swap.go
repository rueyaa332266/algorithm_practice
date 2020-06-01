package main

import "fmt"

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	fmt.Println(swap(1, 2))
}
