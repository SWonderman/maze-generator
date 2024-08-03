package main

import "testing"

type TestNode struct {
	cost float64
}

func (testNode TestNode) GetCost() float64 {
	return testNode.cost
}

func TestPushToEmptyHeap(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})

	if minHeap.Len() != 1 {
		t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 7 {
		t.Fatalf("Unexpected root of the heap. Expected 7 but got %d", (*minHeap)[0])
	}
}

func TestPushTwoElements(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})
	minHeap.Push(TestNode{5})

	if minHeap.Len() != 2 {
		t.Fatalf("Unexpected length of the heap. Expected 2 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 5 {
		t.Fatalf("Unexpected root of the heap. Expected 5 but got %d", (*minHeap)[0])
	}

	expectedList := []int{5, 7}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}
}

func TestPushThreeElements(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})
	minHeap.Push(TestNode{5})
	minHeap.Push(TestNode{9})

	if minHeap.Len() != 3 {
		t.Fatalf("Unexpected length of the heap. Expected 3 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 5 {
		t.Fatalf("Unexpected root of the heap. Expected 5 but got %d", (*minHeap)[0])
	}

	expectedList := []int{5, 7, 9}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}
}

func TestPopHeapOneElement(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})

	if minHeap.Len() != 1 {
		t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 7 {
		t.Fatalf("Unexpected root of the heap. Expected 7 but got %d", (*minHeap)[0])
	}

	expectedList := []int{7}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}

	popedNode := minHeap.Pop()
	if popedNode.GetCost() != 7 {
		t.Fatalf("Unexpected poped heap value. Expected 7 but got %f", popedNode.GetCost())
	}

	if minHeap.Len() != 0 {
		t.Fatalf("Unexpected length of the heap. Expected 0 but got %d", minHeap.Len())
	}
}

func TestPopHeapTwoElements(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})
	minHeap.Push(TestNode{4})

	if minHeap.Len() != 2 {
		t.Fatalf("Unexpected length of the heap. Expected 2 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 4 {
		t.Fatalf("Unexpected root of the heap. Expected 4 but got %d", (*minHeap)[0])
	}

	expectedList := []int{4, 7}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}

	popedNode := minHeap.Pop()
	if popedNode.GetCost() != 4 {
		t.Fatalf("Unexpected poped heap value. Expected 4 but got %f", popedNode.GetCost())
	}

	if minHeap.Len() != 1 {
		t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
	}

	expectedList = []int{7}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}
}

func TestPopHeapThreeElements(t *testing.T) {
	minHeap := &MinHeap{}
	minHeap.Push(TestNode{7})
	minHeap.Push(TestNode{4})
	minHeap.Push(TestNode{8})

	if minHeap.Len() != 3 {
		t.Fatalf("Unexpected length of the heap. Expected 3 but got %d", minHeap.Len())
	}

	if (*minHeap)[0].GetCost() != 4 {
		t.Fatalf("Unexpected root of the heap. Expected 4 but got %d", (*minHeap)[0])
	}

	expectedList := []int{4, 7, 8}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}

	popedNode := minHeap.Pop()
	if popedNode.GetCost() != 4 {
		t.Fatalf("Unexpected poped heap value. Expected 4 but got %f", popedNode.GetCost())
	}

	if minHeap.Len() != 2 {
		t.Fatalf("Unexpected length of the heap. Expected 1 but got %d", minHeap.Len())
	}

	expectedList = []int{7, 8}
	for idx := range expectedList {
		if int((*minHeap)[idx].GetCost()) != expectedList[idx] {
			t.Fatalf("Unexpected heap value at index %d. Expected %d but got %d", idx, expectedList[idx], (*minHeap)[idx])
		}
	}
}
