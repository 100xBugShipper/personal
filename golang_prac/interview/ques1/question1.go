package main

import "fmt"

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		}else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]
	leftSorted := merge_sort(left)
	rightSorted := merge_sort(right)

	return merge(leftSorted, rightSorted)
}


func main() {
	arr := []int{1,5,3,8,5,3}
	fmt.Println(merge_sort(arr))
}
