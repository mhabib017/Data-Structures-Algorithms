package main

import "fmt"

func main() {
	ar := []int{20, 32, 34, 2, 3, 45, 2, 55, 23, 10}
	fmt.Println(ar)

	for i:=1; i< len(ar); i++{
		key:= ar[i]
		j := i-1
		for j>=0 && ar[j]>key {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = key
	}
	fmt.Println(ar)

}