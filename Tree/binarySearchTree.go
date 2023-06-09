package main

import (
	"fmt"
	"math"
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
func (b *binarytree) closestvaluefinder(data int) {
	closestvalue(b.root, data, math.MaxInt32, 0)
}
func closestvalue(root *node, data int, mindiff int, value int) {
	if root == nil {
		fmt.Println("The Closes value is : ", value)
		return
	}
	if root.value < data {
		currentdiff := data - root.value
		if currentdiff < mindiff {
			mindiff = currentdiff
			value = root.value
		}
		closestvalue(root.right, data, mindiff, value)
	} else {
		currentdiff := root.value - data
		if currentdiff < mindiff {
			mindiff = currentdiff
			value = root.value
		}
		closestvalue(root.left, data, mindiff, value)
	}

}
func (b *binarytree) max(root *node, value int) {
	if root == nil {
		fmt.Println(value)
		return
	}
	b.max(root.right, root.value)
}
func (b *binarytree) min(root *node) {
	var value int
	for root != nil {
		value = root.value
		root = root.left
	}
	fmt.Println("The minimum Value = ", value)
}
func searchnode(data int, root *node) (int, *node) {
	if root == nil {
		return 0, nil
	}
	if root.value == data {
		return root.value, root
	}
	if root.value < data {
		return searchnode(data, root.right)
	} else {
		return searchnode(data, root.left)
	}

}

func minValue(root *node) *node {
	if root.left == nil {
		return root
	}
	return minValue(root.left)
}

func (b *binarytree) deletenode(root *node, target int, parent *node) bool {
	if root == nil {
		return false
	}
	loop := root
	found := false
	for loop != nil {
		if loop.value == target {
			found = true
			break
		} else if target < loop.value {
			parent = loop
			loop = loop.left
		} else {
			parent = loop
			loop = loop.right
		}
	}
	if !found {
		return false
	}

	//case 1 for leaf node
	if loop.left == nil && loop.right == nil {
		fmt.Println("Deleting leaf node:", loop.value)
		if parent.left == loop {
			parent.left = nil
		} else {
			parent.right = nil
		}
	}

	//case 2 for one child node
	if loop.left == nil {
		fmt.Println("Deleting node with one right child:", loop.value)
		if parent.left == loop {
			parent.left = loop.right
		} else {
			parent.right = loop.right
		}
	} else if loop.right == nil {
		fmt.Println("Deleting node with one left child:", loop.value)
		if parent.left == loop {
			parent.left = loop.left
		} else {
			parent.right = loop.left
		}
	}

	//case 3 with two child nodes
	if loop.left != nil && loop.right != nil {
		fmt.Println("Deleting node with two children:", loop.value)
		min := minValue(loop.right)
		fmt.Println("Replacing with minimum value in right subtree:", min.value)
		loop.value = min.value
		b.deletenode(loop.right, min.value, loop)
	}

	return true
}

func isBinarySearchTree(root *node, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.value <= min || root.value >= max {
		return false
	}
	return isBinarySearchTree(root.left, min, root.value) && isBinarySearchTree(root.right, root.value, max)

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
	bi.valueinsert(25)
	// bi.valueinsert(21)
	// fmt.Println(bi.count)
	// fmt.Println(bi)
	// fmt.Println(bi.search(45))
	// bi.preorderTraversal(bi.root)
	// bi.inorderTraversal(bi.root)
	// bi.postorderTraversal(bi.root)
	// levelOrder(bi.root)
	bi.closestvaluefinder(13)
	// bi.max(bi.root, 0)
	// bi.min(bi.root)
	// bi.printTree(bi.root, "", true)
	// fmt.Println(bi.deletenode(bi.root, 13, nil))
	// bi.printTree(bi.root, "", true)
	// bi.root.value = 3
	// fmt.Println(bi.root)/
	// bi.printTree(bi.root, "", true)
	// fmt.Println(isBinarySearchTree(bi.root, math.MinInt32, math.MaxInt32))

}
