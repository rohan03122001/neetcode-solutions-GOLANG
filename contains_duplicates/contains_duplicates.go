package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 1}
	// arr2 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicates(arr1))
}

func containsDuplicates(nums []int) bool {

	//MAP{
	//	1: false
	//	2: false
	//}
	numbers := make(map[int]int)

	for _, value := range nums {
		if _, ok := numbers[value]; !ok {
			numbers[value] = value
		} else {
			return true
		}
	}
	return false
}