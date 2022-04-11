package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(a int){
	s.items = append(s.items, a)
}
func (s *stack) pop() int {
	l := len(s.items)-1
	val := s.items[l]
	s.items = s.items[:l]
	return val
}

func main(){
	myStack := stack{}
	myStack.push(10)
	myStack.push(20)
	myStack.push(30)
	fmt.Println(myStack)
	fmt.Println(myStack.pop())
	fmt.Println(myStack)
}