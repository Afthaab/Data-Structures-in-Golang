package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) AddNode(val int) {
	node := &Node{value: val}
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(i, j int) {
	node1 := g.nodes[i]
	node2 := g.nodes[j]
	for node1.next != nil {
		node1 = node1.next
	}
	node1.next = &Node{value: j}
	for node2.next != nil {
		node2 = node2.next
	}
	node2.next = &Node{value: i}
}

func (g *Graph) Search(val int) bool {
	for _, node := range g.nodes {
		for node != nil {
			if node.value == val {
				return true
			}
			node = node.next
		}
	}
	return false
}

func main() {
	graph := &Graph{}
	graph.AddNode(0)
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	fmt.Println(graph.Search(2)) // true
	fmt.Println(graph.Search(3)) // false
}
