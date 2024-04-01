package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	var countNums = make(map[int]int)
	for _, val := range nums {
		countNums[val]++
	}
	if len(countNums) != len(nums) {
		return true
	}
	return false
}
