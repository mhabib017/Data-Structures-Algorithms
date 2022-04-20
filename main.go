package main

import "fmt"

func binarySearch( item, low,high int, arr ...int) int {
	if high <= low {
		if item > arr[low] {
			return low+1
		}else {
			return low
		}
	}

	mid := (low + high)/2;
	if item == arr[mid]  {
		return mid+1;
	}
	if item > arr[mid] {
		return binarySearch(item, mid+1, high, arr...);
	}
	return binarySearch(item, low, mid-1, arr...);
}

func main() {
	ar := []int{20, 32, 34, 2, 3, 45, 2, 55, 23, 10}
	fmt.Println(ar)

	for i:=1; i< len(ar); i++{
		key:= ar[i]
		j := i-1
		loc:= binarySearch(key, 0, j, ar...)
		for j>=loc {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1]= key
	}
	fmt.Println(ar)
}
