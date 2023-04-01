package main

import "fmt"

const stackIndex = 5

var top = -1

var stack [stackIndex]int

func isEmpty() bool {
	if top == -1 {
		return true
	} else {
		return false
	}
}
func isFull() bool {
	if top == stackIndex-1 {
		return true
	} else {
		return false
	}
}
func push(value int) {
	if isFull() == true {
		fmt.Println("Stack is Overflow")
		return
	}
	top++
	stack[top] = value

}
func pop() {
	if isEmpty() == true {
		fmt.Println("Stack is UnderFlow")
		return
	}
	fmt.Println("The Popped Element is: ", stack[top])
	top--

}
func printStack() {
	if isEmpty() == true {
		fmt.Println("Stack is Empty")
		return
	}
	for i := top; i >= 0; i-- {
		fmt.Println(stack[i])
	}
}
func peek() {
	if isEmpty() == true {
		fmt.Println("Stack is Empty")
		return
	}
	fmt.Println(stack[top])
}

func main() {
	push(10)
	push(80)
	push(100)
	push(45)
	pop()
	pop()
	pop()
	pop()
	printStack()
}
