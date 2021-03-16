package main

import "fmt"

type TrieNode struct {
	path int                // 多少个单词公用这个节点
	end  int                // 多少单词以这个节点结尾
	next map[byte]*TrieNode // 下一个字符的节点
}

type TrieTree struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{next: make(map[byte]*TrieNode)}
}

func NewTrieTree() *TrieTree {
	return &TrieTree{NewTrieNode()}
}

func (t *TrieTree) insert(word string) {
	node := t.root
	for i := range word {
		if node.next[word[i]] == nil {
			node.next[word[i]] = NewTrieNode()
		}
		node = node.next[word[i]]
		node.path++
	}
	node.end++ // last char
}

func (t *TrieTree) search(word string) bool {
	node := t.root
	for i := range word {
		if node.next[word[i]] == nil {
			return false
		}
		node = node.next[word[i]]
	}
	return node.end > 0
}

func (t *TrieTree) delete(word string) {
	node := t.root
	if !t.search(word) { // 如果word不在trie tree中，直接返回
		return
	}
	for i := range word {
		// attention!
		node.next[word[i]].path--
		if node.next[word[i]].path == 0 { // 如果下一个node的path为0，说没有单词经过这个节点，直接将node置为nil
			node.next[word[i]] = nil
			return
		}
		node = node.next[word[i]] // 遍历每个word的节点
	}
	node.end--
}

func (t *TrieTree) prefixNum(prefix string) int {
	node := t.root
	for i := range prefix {
		if node.next[prefix[i]] == nil {
			return 0
		}
		node = node.next[prefix[i]]
	}
	return node.path
}

func (t *TrieTree) match(pattern string) bool {
	return t._match(t.root, pattern, 0)
}

// 从node开始，匹配pattern[index...]
// 一个个node和pattern[index]比较
func (t *TrieTree) _match(node *TrieNode, pattern string, index int) bool {
	if index == len(pattern) {
		return node.end > 0
	}

	c := pattern[index]
	if c != '.' {
		return node.next[c] != nil && t._match(node.next[c], pattern, index+1)
	} else {
		// c == '.'
		// iterate all node.next
		for i := range node.next {
			if t._match(node.next[i], pattern, index+1) {
				return true
			}
		}
		return false
	}
}

func main() {
	tree := NewTrieTree()
	tree.insert("abc")
	tree.insert("abcd")
	tree.insert("abd")
	tree.insert("b")
	tree.insert("bcd")
	tree.insert("efg")
	tree.insert("hik")
	fmt.Println(tree.search("abd"))
	fmt.Println(tree.search("abde"))
	fmt.Println(tree.prefixNum("ab"))
	tree.delete("abd")
	fmt.Println(tree.search("abd"))
	fmt.Println(tree.prefixNum("ab"))
	fmt.Println(tree.match("h.."))
	fmt.Println(tree.match("h."))
}
