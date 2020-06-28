package main

import "fmt"

func heapConstructio(heap []int) []int {
	length := len(heap)
	start := parent(length - 1)
	for i := start; i >= 0; i-- {
		heap = downHeap(heap, i)
	}
	return heap
}

func downHeap(heap []int, index int) []int {
	var largest int
	last := len(heap) - 1
	l, r := child(index)
	if l <= last && heap[l] > heap[index] {
		largest = l
	} else {
		largest = index
	}
	if r <= last && heap[r] > heap[largest] {
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

func parent(in int) int {
	if in == 0 {
		return -1
	}
	if in%2 == 0 {
		return in/2 - 1
	}
	return in / 2
}

func child(in int) (int, int) {
	return in*2 + 1, (in + 1) * 2
}

func main() {
	heap := []int{5, 8, 2, 20, 1, 3, 9, 6, 22, 15, 8, 10}
	fmt.Println(heapConstructio(heap))

}
