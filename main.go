package main

import "fmt"

func main() {
	ar := []int{20, 32, 34, 2, 3, 45, 2, 55, 23, 10}

	fmt.Println(ar)

	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	fmt.Println(ar)
}