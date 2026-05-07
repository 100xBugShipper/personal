package main

import "fmt"

func checkSubarray(s1, s2 []int) bool {

	if (len(s1) == len(s2)) && (s1[0] == s2[0]) {
		return true
	}

	s1hash := make(map[int]int)

	for i, v := range s1 {
		s1hash[v] = i
	}

	for _, val := range s2 {
		if _, ok := s1hash[val]; !ok {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 8}
	fmt.Println(checkSubarray(arr1, arr2))
}


// func checkSubarray(s1, s2 []int) bool {
// 	if len(s2) == 0 || len(s2) > len(s1) {
// 		return false
// 	}
//
// 	for i := 0; i <= len(s1)-len(s2); i++ {
// 		match := true
// 		for j := 0; j < len(s2); j++ {
// 			if s1[i+j] != s2[j] {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			return true
// 		}
// 	}
// 	return false
// }
