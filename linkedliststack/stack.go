package main

import "fmt"

type stack struct {
	top    *node
	length int
}

type node struct {
	value int
	next  *node
}

func (l *stack) push(data int) {
	newNode := node{value: data}
	if l.isEmpty() == true {
		l.top = &newNode
	} else {
		newNode.next = l.top
		l.top = &newNode
	}
	l.length++
}
func (l *stack) isEmpty() bool {
	if l.top == nil {
		return true
	} else {
		return false
	}
}
func (l *stack) pop() {
	if l.isEmpty() == true {
		fmt.Println("The stack is already empty")
		return
	}
	l.top = l.top.next
}
func (l *stack) peek() {
	if l.isEmpty() == true {
		fmt.Println("The Top is: ", -1)
		return
	}
	fmt.Println("The Top is: ", l.top.value)
}
func (l *stack) printStack() {
	if l.isEmpty() == true {
		fmt.Println("The stack is empty")
		return
	}
	l.peek()
	fmt.Println("The Stack")
	current := l.top
	for current != nil {
		fmt.Println(current.value, current.next)
		current = current.next
	}
}
func main() {
	li := stack{}
	li.push(10)
	li.push(20)
	li.push(30)
	li.pop()
	li.pop()
	li.pop()
	li.peek()
}
