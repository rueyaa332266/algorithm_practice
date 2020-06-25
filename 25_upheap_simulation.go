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

func main() {
	heap := []int{23, 18, 11, 7, 5, 6, 9, 3, 2, 1, 2, 3}
	heap[9] = 25
	fmt.Println(upHeap(heap, 9))

}
