package main

import "fmt"

func rotate(nums []int, k int) {

	k = k % len(nums)

	// Переворачиваем весь массив -> [ 7 6 5 4 3 2 1 ]
	reverse(nums, 0, len(nums)-1)

	// Переворачиваем первую часть массива - > [ 5 6 7 4 3 2 1 ]
	reverse(nums, 0, k-1)

	// Переворачиваем вторую часть массива -> [ 5 6 7 1 2 3 4 ]
	reverse(nums, k, len(nums)-1)

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	var k = 3

	rotate(nums, k)

	fmt.Println(nums)
}
