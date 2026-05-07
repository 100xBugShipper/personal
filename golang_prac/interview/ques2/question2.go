package main

import "fmt"


func dnf(arr []int) []int {

	low, mid := 0, 0
	high := len(arr) - 1

	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		}else if arr[mid] == 1 {
			mid ++
		}else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high --
		}
	}

	return arr
}

func main() {
	arr := []int{1,0,1,2}
	fmt.Println(dnf(arr))
}
