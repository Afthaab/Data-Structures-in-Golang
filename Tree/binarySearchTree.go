package main

import (
	"fmt"
	// "math"
)

type binarytree struct {
	root  *node
	count int
}

type node struct {
	value int
	left  *node
	right *node
}

func (b *binarytree) valueinsert(data int) {
	b.insert(b.root, data)
}

func (b *binarytree) insert(root *node, data int) {
	newNode := node{value: data}
	if b.root == nil {
		b.root = &newNode
		b.count++
		return
	}
	if data < root.value {
		if root.left == nil {
			root.left = &newNode
			b.count++
			return
		}
		b.insert(root.left, data)
	} else {
		if root.right == nil {
			root.right = &newNode
			b.count++
			return
		}
		b.insert(root.right, data)
	}
}

func (b *binarytree) search(value int) bool {
	return b.searchHelper(b.root, value)
}

func (b *binarytree) searchHelper(current *node, data int) bool {
	if current == nil {
		return false
	}
	if data == current.value {
		return true
	}
	if data < current.value {
		return b.searchHelper(current.left, data)
	} else {
		return b.searchHelper(current.right, data)
	}
}

func (b *binarytree) inorderTraversal(root *node) { //acending order left root right
	if root != nil {
		b.inorderTraversal(root.left)
		fmt.Println(root.value)
		b.inorderTraversal(root.right)
	}
}
func (b *binarytree) preorderTraversal(root *node) { //root left right
	if root != nil {
		fmt.Println(root.value)
		b.preorderTraversal(root.left)
		b.preorderTraversal(root.right)
	}
}
func (b *binarytree) postorderTraversal(root *node) {
	if root != nil {
		b.preorderTraversal(root.left)
		b.preorderTraversal(root.right)
		fmt.Println(root.value)
	}
}

func (b *binarytree) printTree(node *node, indent string, last bool) {
	if node == nil {
		return
	}

	fmt.Print(indent)
	if last {
		fmt.Print("└─")
		indent += "  "
	} else {
		fmt.Print("├─")
		indent += "│ "
	}
	fmt.Println(node.value)

	b.printTree(node.left, indent, node.right == nil)
	b.printTree(node.right, indent, true)
}

func levelOrder(root *node) {
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		fmt.Println(curr.value)
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}
func main() {
	bi := binarytree{}
	bi.valueinsert(13)
	bi.valueinsert(10)
	bi.valueinsert(27)
	bi.valueinsert(7)
	bi.valueinsert(12)
	bi.valueinsert(11)
	bi.valueinsert(23)
	bi.valueinsert(28)
	bi.valueinsert(22)
	bi.valueinsert(24)
	bi.valueinsert(26)
	bi.valueinsert(30)
	// fmt.Println(bi.count)
	// fmt.Println(bi)
	// bi.printTree(bi.root, "", true)
	// fmt.Println(bi.search(45))
	// bi.preorderTraversal(bi.root)
	// bi.inorderTraversal(bi.root)
	// bi.postorderTraversal(bi.root)
	// levelOrder(bi.root)
	// bi.closestvaluefinder(50)
	// bi.max(bi.root, 0)
	// bi.min(bi.root)
}
