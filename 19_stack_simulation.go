package main

import (
	"fmt"
)

type stack []int

func (s *stack) push(in int) {
	*s = append(*s, in)
}

func (s *stack) pop() {
	slice := *s
	*s = slice[:len(*s)-1]
}

func main() {
	var example stack
	example.push(1)
	example.push(2)
	example.push(3)
	example.pop()
	example.pop()
	fmt.Println(example)
}
