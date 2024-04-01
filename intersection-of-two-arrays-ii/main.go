package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var countNums1 = make(map[int]int)
	var countNums2 = make(map[int]int)
	var res []int
	for _, val := range nums1 {
		countNums1[val]++
	}
	for _, val := range nums2 {
		countNums2[val]++
	}
	for key, val := range countNums1 {
		if countNums2[key] != 0 {
			for i := 0; i < val && i < countNums2[key]; i++ {
				res = append(res, key)
			}
		}
	}
	return res
}

func main() {
	var nums1 = []int{4, 9, 5}
	var nums2 = []int{9, 4, 9, 8, 4}
	res := intersect(nums1, nums2)
	fmt.Println(res)
}
