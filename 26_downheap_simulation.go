package main

import "fmt"

func downHeap(heap []int, index int) []int {
	var largest int
	last := len(heap) - 1
	l, r := child(index)
	if l < last && heap[l] > heap[index] {
		largest = l
	} else {
		largest = index
	}
	if r < last && heap[r] > heap[largest] {
		largest = r
	}

	if largest != index {
		tmp := heap[index]
		heap[index] = heap[largest]
		heap[largest] = tmp
		downHeap(heap, largest)
	}
	return heap
}

func child(in int) (int, int) {
	return in*2 + 1, (in + 1) * 2
}

func main() {
	heap := []int{20, 18, 11, 5, 7, 6, 9, 3, 2, 4, 2, 5}
	heap[0] = 1
	fmt.Println(downHeap(heap, 0))

}
