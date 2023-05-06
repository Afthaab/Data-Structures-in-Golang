package main

import "fmt"

type trieNode struct {
	children map[rune]*trieNode
	isEnd    bool
}

type trieDS struct {
	root *trieNode
}

func (t *trieDS) insert(word string) {
	current := t.root
	for _, char := range word {
		childnode, ok := current.children[char]
		if !ok {
			childnode = &trieNode{children: make(map[rune]*trieNode), isEnd: false}
			current.children[char] = childnode
		}
		current = childnode
	}
	current.isEnd = true
}

func (t *trieDS) search(word string) bool {
	current := t.root
	for _, char := range word {
		childNode, ok := current.children[char]
		if !ok {
			return false
		}
		current = childNode
	}
	return current.isEnd
}

func (t *trieDS) Delete(word string) bool {
	currentNode := t.root
	var prevNodes []*trieNode // Track nodes leading up to the last character of the word
	for _, char := range word {
		childNode, ok := currentNode.children[char]
		if !ok {
			return false // Word doesn't exist in the Trie
		}
		prevNodes = append(prevNodes, currentNode)
		currentNode = childNode
	}

	if !currentNode.isEnd {
		return false // Word doesn't exist in the Trie
	}

	currentNode.isEnd = false // Unmark the last node as the end of a word

	// Traverse back up the Trie, removing any nodes that no longer lead to the end of a word
	for i := len(prevNodes) - 1; i >= 0; i-- {
		node := prevNodes[i]
		if len(node.children) == 0 && !node.isEnd {
			delete(node.children, 0)
		} else {
			break
		}
	}
	return true // Word was successfully deleted from the Trie
}

func (t *trieDS) searchPrefix(prefix string) []string {
	current := t.root
	for _, char := range prefix {
		childnode, ok := current.children[char]
		if !ok {
			return []string{}
		}
		current = childnode
	}
	return t.searchWords(current, prefix)
}

func (t *trieDS) searchWords(node *trieNode, prefix string) []string {
	result := []string{}
	if node.isEnd {
		result = append(result, prefix)
	}
	for char, childnode := range node.children {
		childWords := t.searchWords(childnode, prefix+string(char))
		result = append(result, childWords...)
	}
	return result
}

func (t *trieDS) Print() {
	t.root.printHelper("")
}

func (n *trieNode) printHelper(prefix string) {
	if n.isEnd {
		fmt.Println(prefix)
	}
	for char, child := range n.children {
		child.printHelper(prefix + string(char))
	}
}

func main() {
	trie := &trieDS{root: &trieNode{children: make(map[rune]*trieNode), isEnd: false}}
	trie.insert("afthab")
	trie.insert("bantu")
	trie.insert("banti")
	// trie.Delete("bantu")
	fmt.Println(trie.searchPrefix("ba"))
	// trie.Print()
}
