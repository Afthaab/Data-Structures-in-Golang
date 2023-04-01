package main

import "fmt"

type queue struct {
	front *node
	rear  *node
}

type node struct {
	value int
	next  *node
}

func (l *queue) isEmpty() bool {
	if l.front == nil && l.rear == nil {
		return true
	} else {
		return false
	}
}
func (l *queue) enqueue(data int) {
	newNode := node{value: data}
	if l.isEmpty() == true {
		l.front = &newNode
		l.rear = &newNode
	} else {
		l.rear.next, l.rear = &newNode, &newNode
	}
}
func (l *queue) peek() {
	if l.isEmpty() == true {
		fmt.Println("The Front and Rear is", -1, -1)
		return
	}
	fmt.Println("The Front: ", l.front.value)
	fmt.Println("The Rear: ", l.rear.value)
}
func (l *queue) printLinkedlist() {
	if l.isEmpty() == true {
		fmt.Println("The queue is empty")
		return
	}
	current := l.front
	fmt.Println("The Queue")
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func (l *queue) dequeue() {
	if l.isEmpty() == true {
		fmt.Println("The queue already is empty")
		return
	}
	if l.front == l.rear {
		l.rear = nil
		l.front = nil
	} else {
		l.front = l.front.next
	}
}

func main() {
	q := queue{}
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.printLinkedlist()
	q.peek()
}
