package main

import (
	"fmt"
)

type queue struct {
	que   []int
	first int
	last  int
}

func (s *queue) enqueue(in int) {
	s.que[s.last] = in
	s.last = s.last + 1
}

func (s *queue) dequeue() {
	s.que[s.first] = 0
	s.first = s.first + 1
}

func main() {
	var que = make([]int, 6)
	example := queue{que, 0, 0}
	example.enqueue(4)
	example.enqueue(8)
	example.enqueue(1)
	example.dequeue()
	example.enqueue(7)
	example.dequeue()
	fmt.Println(example.que)
}
