package main

import (
	"fmt"
)

// index is the size of the hashtable which is the size of the array
const Index = 7

// hash table structure : it will hold an array
type Hashtable struct {
	array [Index]*bucket
}

// bucket structure : bucket is the linked list in each slot in the array
type bucket struct {
	head *bucketnode
}

// bucket node structure: it is the node in the linkedlist that holds the key
type bucketnode struct {
	key  string
	next *bucketnode
}

// init function: this function is used to initialize a bucket to the each slot in the linked list
func Initt() *Hashtable {
	result := &Hashtable{}
	for i := 0; i < Index; i++ {
		result.array[i] = &bucket{}
	}
	return result
}

// hash table operation
// insert function(): This func will take the key and store it in the hash table array
func (h *Hashtable) insert(key string) {
	index := hashfunc(key)
	h.array[index].insert(key)
}

// search func(): This will search whether the key is present in the hash table and return a bool value
func (h *Hashtable) search(key string) bool {
	index := hashfunc(key)
	return h.array[index].search(key)
}

// delete func(): This will delete a key from the hash table
func (h *Hashtable) delete(key string) {
	index := hashfunc(key)
	h.array[index].delete(key)
}

// bucket operation
// insert the key to the bucket node
func (b *bucket) insert(k string) {
	newBucket := &bucketnode{key: k}
	newBucket.next, b.head = b.head, newBucket
}

// search for a key in the bucket
func (b *bucket) search(k string) bool {
	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}

// delete a key from the bucket
func (b *bucket) delete(k string) {
	if b.head == nil {
		return
	}

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prev := b.head
	current := b.head.next

	for current != nil {
		if current.key == k {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

// hash func which will convert string to ascii and finds the value
func hashfunc(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Index
}

func main() {
	add := Initt()
	fmt.Println(add)

	add.insert("RANDY")
	fmt.Println(add.search("RANDY"))

	add.delete("RANDY")
	fmt.Println(add.search("RANDY"))
}
