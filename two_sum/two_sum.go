package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 1}
	// arr2 := []int{1, 2, 3, 4}
	fmt.Println(twoSum(arr1, 4))
}

func twoSum(nums []int, target int) []int {

	numbers := make(map[int]int)

	for index, value := range nums {
		complement := target - value
		
		if result, ok := numbers[complement]; !ok {

			numbers[value] = index

		} else{


			return []int{index, result}
		}
	}
	return []int{-1, -1}
}