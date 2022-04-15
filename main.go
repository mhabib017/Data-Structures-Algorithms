package main

import (
	"fmt"
)

const ARRAY_SIZE = 5

type HashTable struct {
	array [ARRAY_SIZE]*Bucket
}

func (h *HashTable) insert(key string){
	index:= hash(key)
	h.array[index].insert(key)
}

func (h HashTable) search(key string)bool {
	index:= hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) delete(key string)  {
	index:= hash(key)
	h.array[index].delete(key)
}


type Bucket struct {
	head *BucketNode
}

type BucketNode struct{
	key string
	next *BucketNode
}

func (b *Bucket) insert(key string){
	
	newNode := BucketNode{key:key}
	newNode.next = b.head
	b.head = &newNode
	
}

func (b *Bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousToDelete := b.head
	for previousToDelete.next != nil {
		if previousToDelete.next.key == key{
			previousToDelete.next = previousToDelete.next.next
			return
		}
		previousToDelete = previousToDelete.next
	}
}

func (b Bucket) search(key string) bool{
	current:= b.head
	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false
}


func initHashTable() *HashTable{
	hashTable:= HashTable{}
	for i:= range  hashTable.array{
		hashTable.array[i] = &Bucket{}
	}
	return &hashTable
}

func hash(key string) int{
	sum:=0
	for i:= range key{
		sum+= int(i)
	}
	return sum % ARRAY_SIZE
}

func main() {
	hashTable := initHashTable()

	hashTable.insert("Randy")
	hashTable.insert("Randy")
	hashTable.insert("Eric")
	fmt.Println(hashTable.search("Randy"))
	hashTable.delete("Randy")
	fmt.Println(hashTable.search("Randy"))
	fmt.Println(hashTable.search("Ben"))

}