package main

import "fmt"

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for i = 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
