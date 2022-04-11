package main

import "fmt"

func binarySearch(v int,ar ...int) bool {
	
	low:=0
	high:= len(ar)-1
	
	for low<=high {
		mid:= (low+high)/2
	
		if ar[mid] == v {
			return true
		} else if ar[mid] < v {
			low = mid+1
		}else if ar[mid]> v {
			high = mid -1
		}
		
		
	}
	fmt.Println(ar)
	return false
}

func main(){
	ar := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println( binarySearch(70, ar...))
}