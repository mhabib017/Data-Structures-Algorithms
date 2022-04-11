package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}

func (n *Node) insert(val int) {

	if n.key < val {
		if n.right == nil {
			n.right = &Node{key: val}
		} else {
			n.right.insert(val)
		}
	} else if n.key > val {
		if n.left == nil {
			n.left = &Node{key: val}
		} else {
			n.left.insert(val)
		}
	}
}

func (n *Node) inOrderTraversal(){
	if n!= nil {
		n.left.inOrderTraversal();
		fmt.Print(n.key, " ")
		n.right.inOrderTraversal();
	}
}

func (n *Node) preOrderTraversal(){
	if n!= nil {
		fmt.Print(n.key, " ")
		n.left.inOrderTraversal();
		n.right.inOrderTraversal();
	}
}

func (n *Node) postOrderTraversal(){
	if n!= nil {
		n.left.inOrderTraversal();
		n.right.inOrderTraversal();
		fmt.Print(n.key, " ")
	}
}


func main() {
	tree := Node{key: 50}
	tree.insert(20)
	tree.insert(40)
	tree.insert(10)
	tree.insert(11)
	tree.inOrderTraversal()
	fmt.Println()
	tree.preOrderTraversal()
	fmt.Println()
	tree.postOrderTraversal()
}