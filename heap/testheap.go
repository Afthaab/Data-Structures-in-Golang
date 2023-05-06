package main

import "fmt"

type heap struct {
	arr  []int
	size int
}

func parentNode(i int) int {
	return (i - 1) / 2
}

func (h *heap) insertheap(data int) {
	if len(h.arr) == h.size {
		fmt.Println("the heap size is full")
		return
	}
	h.size++
	i := h.size - 1
	h.arr[i] = data
	for i != 0 && h.arr[parentNode(i)] < h.arr[i] {
		h.arr[parentNode(i)], h.arr[i] = h.arr[i], h.arr[parentNode(i)]
		i = parentNode(i)
	}
}

func (h *heap) deleteNode() {
	if len(h.arr) == 0 {
		fmt.Println("the heap is empty")
		return
	}
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.heapify(0)
}
func (h *heap) heapify(i int) { //recursion
	larger := i
	left := 2*i + 1
	ritgh := 2*i + 2
	if left < h.size && h.arr[left] > h.arr[larger] {
		larger = left
	}
	if ritgh < h.size && h.arr[ritgh] > h.arr[larger] {
		larger = ritgh
	}
	if larger != i {
		h.swap(i, larger)
		h.heapify(larger)
	}
}
func (m *heap) swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}
func heapSort(arr []int) []int {
	h := &heap{arr, len(arr)}
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
	for i := h.size - 1; i >= 0; i-- {
		h.swap(0, i)
		h.size--
		h.heapify(0)
	}
	return h.arr
}

func main() {
	// newHeap := &heap{arr: make([]int, 10), size: 0}
	// newHeap.insertheap(10)
	// newHeap.insertheap(20)
	// newHeap.insertheap(110)
	// newHeap.insertheap(400)
	// newHeap.insertheap(500)
	// newHeap.insertheap(650)
	// newHeap.insertheap(0)
	// newHeap.insertheap(40)
	// newHeap.insertheap(710)
	// newHeap.insertheap(50)
	// fmt.Println(newHeap.arr)
	// newHeap.deleteNode()
	// fmt.Print(newHeap.arr)
	array := []int{7, 5, 6, 10, 85, 4}
	fmt.Println(heapSort(array))

}
