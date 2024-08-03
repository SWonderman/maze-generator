package main

type HeapNode interface {
	GetCost() float64
}

type MinHeap []HeapNode

func (heap MinHeap) swapByIndex(i int, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap MinHeap) leftChildIndex(i int) int {
	return 2*i + 1
}

func (heap MinHeap) rightChildIndex(i int) int {
	return 2*i + 2
}

func (heap MinHeap) parentIndex(i int) int {
	return int((i - 1) / 2)
}

func (heap MinHeap) Len() int {
	return len(heap)
}

func (heap *MinHeap) heapifyUp(i int) {
	parentIdx := heap.parentIndex(i)
	for (*heap)[parentIdx].GetCost() > (*heap)[i].GetCost() {
		heap.swapByIndex(parentIdx, i)
	}
}

func (heap *MinHeap) heapifyDown(i int) {
	leftChildIdx := heap.leftChildIndex(i)
	rightChildIdx := heap.rightChildIndex(i)

	smallerChildIdx := i

	if leftChildIdx <= heap.Len()-1 && (*heap)[smallerChildIdx].GetCost() > (*heap)[leftChildIdx].GetCost() {
		smallerChildIdx = leftChildIdx
	}

	if rightChildIdx <= heap.Len()-1 && (*heap)[smallerChildIdx].GetCost() > (*heap)[rightChildIdx].GetCost() {
		smallerChildIdx = rightChildIdx
	}

	if smallerChildIdx != i {
		heap.swapByIndex(smallerChildIdx, i)
		heap.heapifyDown(smallerChildIdx)
	}
}

func (heap *MinHeap) Push(node HeapNode) {
	*heap = append(*heap, node)

	newValueIdx := heap.Len() - 1
	parent := (*heap)[heap.parentIndex(newValueIdx)]
	if parent.GetCost() > node.GetCost() {
		heap.heapifyUp(newValueIdx)
	}
}

func (heap *MinHeap) Pop() HeapNode {
	if heap.Len() == 0 {
		return nil
	}

	minNode := (*heap)[0]

	// Swap the current root with the last value,
	// discard the last value in the heap,
	// and heapify down to retaing the heap property.

	(*heap)[0] = (*heap)[heap.Len()-1]
	*heap = (*heap)[0 : heap.Len()-1]

	heap.heapifyDown(0)

	return minNode
}
