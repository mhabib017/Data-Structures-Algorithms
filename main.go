package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(a int) {
	q.items = append(q.items, a)
}

func (q *Queue) dequeue() int {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func main(){
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(10)
	myQueue.enqueue(20)
	myQueue.enqueue(30)
	myQueue.enqueue(40)
	fmt.Println(myQueue)
	
	fmt.Println(myQueue.dequeue())
	
	fmt.Println(myQueue)
	
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.dequeue())
	
	fmt.Println(myQueue)

}