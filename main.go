package main

import "fmt"

type Node struct {
	key int
	left *Node
	right *Node
}

func (n *Node) insert(val int) {
	
	if n.key < val {
		if n.right == nil {
			n.right = &Node{key: val}
		}else {
			n.right.insert(val)
		}
	}else if n.key > val {
		if n.left == nil {
			n.left =  &Node{key: val}
		}else {
			n.left.insert(val)
		}
	}
}

func (n *Node) search(val int) bool {
	if n==nil {
		return false
	}
	if n.key < val {
		return n.right.search(val)
	}else if n.key >val {
		return n.left.search(val)
	}
	return true
}

func  main()  {
	tree:= &Node{key: 50}
	tree.insert(20)
	tree.insert(40)
	tree.insert(10)
	tree.insert(11)
	fmt.Println(tree.search(60))
	fmt.Println(tree.search(40))
	fmt.Println(tree.search(110))
	fmt.Println(tree.search(11))
}