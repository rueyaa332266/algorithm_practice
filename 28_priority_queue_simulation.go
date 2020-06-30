package main

import "fmt"

func upHeap(heap []int, index int) []int {
	for {
		parentIndex := parent(index)
		if parentIndex == -1 {
			break
		}
		if heap[index] < heap[parentIndex] {
			break
		}
		if heap[index] > heap[parentIndex] {
			tmp := heap[parentIndex]
			heap[parentIndex] = heap[index]
			heap[index] = tmp
		}
		index = parentIndex
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

func child(in int) (int, int) {
	return in*2 + 1, (in + 1) * 2
}

func insert(heap []int, in int) ([]int, int) {
	last := len(heap) - 1
	for i := 0; i < len(heap); i++ {
		if heap[i] == 0 {
			last = i - 1
			break
		}
	}
	heap[last+1] = in
	return heap, last + 1
}

func extract(heap []int) []int {
	last := len(heap) - 1
	for i := 0; i < len(heap); i++ {
		if heap[i] == 0 {
			last = i - 1
			break
		}
	}
	heap[0] = heap[last]
	heap[last] = 0
	return heap
}

func main() {
	heap := []int{20, 9, 11, 7, 3, 6, 9, 3, 2, 0, 0, 0}
	var last int

	heap, last = insert(heap, 28)
	upHeap(heap, last)
	heap, last = insert(heap, 22)
	upHeap(heap, last)
	heap = extract(heap)
	heap = downHeap(heap, 0)
	heap, last = insert(heap, 27)
	upHeap(heap, last)
	heap = extract(heap)
	heap = downHeap(heap, 0)

	fmt.Println(heap)
}
