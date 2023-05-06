package main

import (
	"fmt"
	"math"
)

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func (t *tree) insert(data int, root *node) { //recursion
	newNode := &node{value: data}
	if t.root == nil {
		t.root = newNode
		return
	}
	if data < root.value {
		if root.left == nil {
			root.left = newNode
			return
		}
		t.insert(data, root.left)
	} else {
		if root.right == nil {
			root.right = newNode
			return
		}
		t.insert(data, root.right)
	}

}

func (t *tree) search(data int, root *node) bool { //recursion
	if root == nil {
		return false
	}
	if data == root.value {
		return true
	}
	if data < root.value {
		return t.search(data, root.left)
	} else {
		return t.search(data, root.right)
	}
}

func (t *tree) inorderTraversal(root *node) { //recursion
	if root != nil {
		t.preorderTraversal(root.left)
		fmt.Println(root.value)
		t.preorderTraversal(root.right)
	}
}
func (t *tree) postorderTraversal(root *node) { //recursion
	if root != nil {
		t.preorderTraversal(root.left)
		t.preorderTraversal(root.right)
		fmt.Println(root.value)
	}
}
func (t *tree) preorderTraversal(root *node) { //recursion
	if root != nil {
		fmt.Println(root.value)
		t.preorderTraversal(root.left)
		t.preorderTraversal(root.right)
	}
}

func (t *tree) levelorder(root *node) {
	if root == nil {
		return
	}
	queue := []*node{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (t *tree) delete(data int, parent *node, root *node) bool {
	if t.root == nil {
		return false
	}
	current := root
	found := false
	for current != nil {
		if current.value == data {
			found = true
			break
		} else if current.value < data {
			parent = current
			current = current.right
		} else if current.value > data {
			parent = current
			current = current.left
		}
	}
	fmt.Println("ji")
	if found == false {
		return false
	}
	fmt.Println("ji")
	fmt.Println("the node:", current.value, parent.value)
	//case 1 delete childnode
	if current.left == nil && current.right == nil {
		if parent.left == current {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return true
	}
	//one child node
	if current.left == nil || current.right == nil {
		if parent.left == current {
			if current.left == nil {
				parent.left = current.right
			} else if current.right == nil {
				parent.left = current.left
			}
			return true
		} else if parent.right == current {
			if current.left == nil {
				parent.right = current.right
			} else if current.right == nil {
				parent.right = current.left
			}
			return true
		}
	}
	//two children
	if current.left != nil && current.right != nil {
		min := minValue(current.right)
		current.value = min.value
		t.delete(min.value, current, current.right)
	}
	return true
}

func minValue(root *node) *node {
	if root != nil {
		root = root.left
	}
	return root
}

func (t *tree) closestvalue(data int, value int, mindiff int, root *node) {
	if root == nil {
		fmt.Println("The closest value is :", value)
		return
	}
	if data < root.value {
		currentdiff := root.value - data
		if currentdiff < mindiff {
			mindiff = currentdiff
			value = root.value
		}
		t.closestvalue(data, value, mindiff, root.left)
	} else {
		currentdiff := data - root.value
		if currentdiff < mindiff {
			mindiff = currentdiff
			value = root.value
		}
		t.closestvalue(data, value, mindiff, root.right)
	}
}

func main() {
	binary := &tree{}
	binary.insert(1100, binary.root)
	binary.insert(200, binary.root)
	binary.insert(1300, binary.root)
	binary.insert(300, binary.root)
	binary.insert(75, binary.root)
	binary.insert(20, binary.root)
	binary.insert(900, binary.root)
	binary.insert(30, binary.root)
	// fmt.Println(binary.search(400, binary.root))
	// binary.preorderTraversal(binary.root)
	// binary.levelorder(binary.root)
	// fmt.Println(binary.delete(30, nil, binary.root))
	// fmt.Println("----------------------")
	// binary.levelorder(binary.root)
	// binary.inorderTraversal(binary.root)
	binary.closestvalue(1101, 0, math.MaxInt32, binary.root)

}
