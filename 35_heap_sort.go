package main

import "fmt"

func heapSort(heap []int) []int {
	heapSize := len(heap)
	for i := heapSize/2 - 1; i >= 0; i-- {
		downHeap(heap, i, 0)
	}
	cut := 1
	for heapSize >= 2 {
		heap[0], heap[heapSize-1] = swap(heap[0], heap[heapSize-1])
		heapSize--
		downHeap(heap, 0, cut)
		cut++
	}
	return heap
}

func downHeap(heap []int, index int, cut int) []int {
	var largest int
	last := len(heap) - 1 - cut
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
		downHeap(heap, largest, cut)
	}
	return heap
}

func child(in int) (int, int) {
	return in*2 + 1, (in + 1) * 2
}

func swap(a int, b int) (int, int) {
	t := a
	a = b
	b = t
	return a, b
}

func main() {
	heap := []int{5, 20, 6, 11, 1, 8, 15, 23, 3}
	fmt.Println(heapSort(heap))
}
