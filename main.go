package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp( len(h.array)-1)
}

func (h *MaxHeap) maxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index-1)/2
}

func left(index int) int {
	return 2*index +1
}

func right(index int) int {
	return 2*index +2
}

func (h *MaxHeap) swap(indexOne, indexTwo int){
	h.array[indexOne], h.array[indexTwo] = h.array[indexTwo], h.array[indexOne]
}

func (h *MaxHeap) extract() int {
	extracted:= h.array[0]
	l:= len(h.array)-1
	if len(h.array) == 0 {
		fmt.Println("can not extract becausue array is empty.")
		return -1
	}
	h.array[0]= h.array[l]
	h.array = h.array[:l]
	
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int){
	lastIndex:= len(h.array) -1
	l, r := left(index), right(index)
	childToCompare:=0

	for l<= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right (index)
		}else {
			return
		}
	}
}


func main(){
	heap:= MaxHeap{}
	ar:= []int{10,20,30, 7,4,3,23,33}
	fmt.Println(heap)
	for _, v:=range ar{
		heap.insert(v)
		fmt.Println(heap)
	}

	for i:=0; i<=5; i++ {
		heap.extract()
		fmt.Println(heap)
	}
}