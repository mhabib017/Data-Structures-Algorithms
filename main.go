package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepand(v int) {
	node := Node{data: v}
	node.next = l.head
	l.head = &node
	l.length++
}

func (l *LinkedList) printList(){
	node:=l.head
	length := l.length
	for length !=0 {
		fmt.Print(node.data, " ")
		node = node.next
		length--
	}
	fmt.Println()
}

func (l *LinkedList) delete(val int) {
	if l.length ==0{
		return 
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	linkedList := LinkedList{}
	linkedList.prepand(10)
	linkedList.prepand(20)
	linkedList.prepand(30)
	linkedList.prepand(40)
	linkedList.prepand(50)
	linkedList.prepand(60)
	linkedList.prepand(70)
	linkedList.prepand(80)
	linkedList.printList()
	linkedList.delete(60)
	linkedList.delete(70)
	linkedList.delete(80)
	linkedList.delete(10)
	linkedList.delete(20)
	linkedList.printList()
}