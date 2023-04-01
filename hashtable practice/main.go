package main

import "fmt"

const Index = 5

type hashtable struct {
	arr [Index]*bucket
}

type bucket struct {
	head *bucketnode
}

type bucketnode struct {
	key   string
	value string
	next  *bucketnode
}

func initt() *hashtable {
	result := &hashtable{}
	for i := 0; i < Index; i++ {
		result.arr[i] = &bucket{}
	}
	return result
}

var top = -1

var stack []int

func push(value int) {
	size := len(stack)
	if top == size-1 {
		fmt.Println("The Stack Full")
		return
	}
	top++
	stack[top] = value
}

func pop() {
	if top < 0 {
		fmt.Println("The Stack is Empty")
		return
	}
	fmt.Println("The Deleted item is:", stack[top])
	top--
}

func main() {
	push(10)
	push(30)
	push(20)
	pop()
	pop()
	fmt.Println(stack)
}

func mergesort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (high + low) / 2
	mergesort(arr, low, mid)
	mergesort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) {
	solo := make([]int, high-low+1)
	left := low
	right := mid + 1
	k := 0
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			solo[k] = arr[left]
			left++
		} else {
			solo[k] = arr[right]
			right++
		}
		k++
	}
	for left <= mid {
		solo[k] = arr[left]
		left++
		k++
	}
	for right <= high {
		solo[k] = arr[right]
		right++
		k++
	}
	for i := 0; i < len(solo); i++ {
		arr[low+i] = solo[i]
	}
}

func quickSort(arr []int, low int, high int) {

}
