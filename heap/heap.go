package main

import "fmt"

type MaxHeap struct {
	arr  []int
	size int
}

func parentNode(i int) int {
	return (i - 1) / 2
}

func (m *MaxHeap) swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *MaxHeap) insertNode(key int) {
	if m.size == len(m.arr) {
		fmt.Println("Heap is full")
		return
	}
	m.size++
	i := m.size - 1
	m.arr[i] = key
	for i != 0 && m.arr[parentNode(i)] < m.arr[i] {
		m.swap(i, parentNode(i))
		i = parentNode(i)
	}

}

func (m *MaxHeap) heapify(i int) {
	larger := i
	left := 2*i + 1
	right := 2*i + 2
	if left < m.size && m.arr[left] > m.arr[larger] {
		larger = left
	}
	if right < m.size && m.arr[right] > m.arr[larger] {
		larger = right
	}
	if larger != i {
		m.swap(i, larger)
		m.heapify(larger)
	}
}

func (m *MaxHeap) deleteNode() {
	if m.size == 0 {
		fmt.Println("heap is empty")
		return
	}
	if m.size == 1 {
		m.size--
		return
	}
	m.arr[0] = m.arr[m.size-1]
	m.size--
	m.heapify(0)
}

func main() {
	heap := &MaxHeap{arr: make([]int, 5), size: 0}
	heap.insertNode(100)
	heap.insertNode(200)
	heap.insertNode(300)
	heap.insertNode(250)
	heap.insertNode(1150)
	fmt.Println(heap)
	heap.deleteNode()
	fmt.Println(heap)

}
