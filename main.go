package main

import "fmt" 

type Node struct {
	children [26]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func InitTie() *Trie {
	trie:= Trie{root: &Node{}}
	return &trie
}

func (t *Trie) insert(w string) {
	wordLength:= len(w)
	currentNode:= t.root
	for i:=0; i< wordLength; i++{
		currentIndex:= w[i]- 'a'
		if currentNode.children[currentIndex] == nil {
			currentNode.children[currentIndex] = &Node{}
		}
		currentNode = currentNode.children[currentIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) search(w string) bool {
	wordLength:= len(w)
	currentNode:= t.root
	for i:=0; i< wordLength; i++{
		currentIndex:= w[i]- 'a'
		if currentNode.children[currentIndex] == nil {
			return false
		}
		currentNode = currentNode.children[currentIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}


func  main()  {
	trie:= InitTie()
	trie.insert("habib")
	trie.insert("happy")
	fmt.Println(trie.search("happy"))
	fmt.Println(trie.search("happyy"))
	fmt.Println(trie.search("happ"))
}