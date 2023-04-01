package main

import "fmt"

var queue [5]int

const queueIndex = 5

var front = -1
var rear = -1

func isEmpty() bool {
	if rear < 0 && front < 0 {
		return true
	} else {
		return false
	}
}
func isFull() bool {
	if rear == queueIndex-1 {
		return true
	} else {
		return false
	}
}
func enqueue(value int) {
	if isFull() == true {
		fmt.Println("Queue is full/ Overloaded")
		return
	}
	if front < 0 && rear < 0 {
		front++
		rear++
		queue[rear] = value

	} else {
		rear++
		queue[rear] = value
	}
}
func dequeue() {
	if isEmpty() == true {
		fmt.Println("Queue is Empty")
		return
	}
	if front == rear {
		fmt.Println("Element Deleted is: ", queue[front])
		front, rear = -1, -1
	} else {
		fmt.Println("Element Deleted is: ", queue[front])
		front++
	}

}
func printQueue() {
	if isEmpty() == true {
		fmt.Println("The Queue is Empty")
		return
	}
	fmt.Println("The Queue is :")
	for i := front; i <= rear; i++ {
		fmt.Println(queue[i])
	}
}
func peek() {
	if rear < 0 && front < 0 {
		fmt.Println("The Rear is :", rear)
		fmt.Println("The Front is :", front)
		return
	}
	fmt.Println("The Rear is :", queue[rear])
	fmt.Println("The Front is :", queue[front])
}

func main() {
	enqueue(10)
	enqueue(20)
	enqueue(30)
	enqueue(40)
	enqueue(50)
	dequeue()
	dequeue()
	dequeue()
	dequeue()
	dequeue()
	printQueue()
	peek()
}
