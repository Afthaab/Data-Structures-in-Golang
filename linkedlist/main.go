package main

import "fmt"

type node struct {
	value int
	next  *node
}
type linkedlist struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedlist) addNode(data int) {
	newNode := node{value: data}
	if l.head == nil {
		l.head = &newNode
		l.length++
	} else {
		l.tail.next = &newNode
		l.length++
	}
	l.tail = &newNode
}
func (l *linkedlist) printLinkedlist() {
	loop := l.head
	fmt.Println("The length: ", l.length)
	fmt.Println("The Head: ", l.head)
	fmt.Println("The Tail: ", l.tail)
	for loop != nil {
		fmt.Println("Value: ", loop.value, "Next Node:", loop.next)
		loop = loop.next
	}
}
func (l *linkedlist) prepend(data int) {
	if l.head == nil {
		fmt.Println("Add a Node to Prepend")
	} else {
		newNode := node{value: data, next: l.head}
		l.head = &newNode
		l.length++
	}
}
func (l *linkedlist) append(data int) {
	newNode := node{value: data}
	if l.tail == nil {
		fmt.Println("There is no tail")
	} else {
		l.tail.next, l.tail = &newNode, &newNode
		l.length++
	}
}
func (l *linkedlist) insertdataAfter(data int, target int) {
	loop := l.head
	if l.head == nil && l.tail == nil {
		fmt.Println("Head and tail are nil")
		return
	}
	if loop.value == target { //best case T[o(1)] s[o(1)]
		newNode := node{value: data, next: loop.next}
		loop.next = &newNode
		l.length++
		return
	}
	if l.tail.value == target { //best case T[o(1)] s[o(1)]
		newNode := node{value: data}
		l.tail.next, l.tail = &newNode, &newNode
		l.length++
		return
	}
	for loop != nil { //worst case T[o(n)] s[o(1)]
		if loop.value == target {
			newNode := node{value: data, next: loop.next}
			loop.next = &newNode
			l.length++
			break
		}
		loop = loop.next
	}
	if loop == nil {
		fmt.Println("Search Value not found in list")
	}

}
func (l *linkedlist) addBefore(data int, target int) {
	var prev *node
	loop := l.head
	if l.head == nil && l.tail == nil {
		fmt.Println("No nodes in the linked list")
		return
	}
	if loop.value == target { //best case T[0(1)] S[o(1)]
		newNode := node{value: data, next: loop}
		l.head = &newNode
		l.length++
		return
	}
	if l.tail.value == target { //best case T[0(1)] S[o(1)]
		newNode := node{value: data}
		l.tail.next, l.tail = &newNode, &newNode
		l.length++
		return
	}
	for loop != nil { // worst case T[o(n)] s[o(1)]
		if loop.value == target {
			newNode := node{value: data, next: loop}
			prev.next = &newNode
			l.length++
			break
		}
		prev = loop
		loop = loop.next
	}
	if loop == nil {
		fmt.Println("The Give value is not found in the linked list")
	}

}
func (l *linkedlist) deleteNode(target int) {
	var prev *node
	loop := l.head
	if l.head == nil && l.tail == nil {
		fmt.Println("No node is present in the linked list")
		return
	}
	if loop.value == target { //best case T[0(1)] S[o(1)]
		l.head = loop.next
		return
	}
	for loop != nil && loop.value != target { // worst case T[o(n)] s[o(1)]
		prev = loop
		loop = loop.next
	}
	if l.tail.value == target { //best case T[0(1)] S[o(1)]
		l.tail = prev
		prev.next = nil
		return
	}
	prev.next = loop.next

}
func (l *linkedlist) deleteOneOfMany(target int) {
	loop := l.head
	var prev *node
	for loop != nil {
		if loop.value == target && loop == l.head {
			l.head = loop.next
			loop = loop.next
			l.length--
		} else if loop.value == target && loop != l.tail {
			prev.next = loop.next
			loop = loop.next
			l.length--
		} else if loop.value == target && loop == l.tail {
			l.tail = prev
		} else {
			prev = loop
			loop = loop.next
		}
	}
}

func main() {
	var mu linkedlist
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	// mu.addNode(20)
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	// mu.addNode(40)
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	mu.addNode(10)
	mu.printLinkedlist()
	mu.deleteOneOfMany(10)
	mu.printLinkedlist()
}
