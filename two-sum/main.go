package main

func TwoSum(nums []int, target int) []int {

	var res = map[int]int{}

	for i, num := range nums {

		if _, ok := res[num]; ok {
			return []int{i, res[num]}
		} else {
			res[target-num] = i
		}
	}
	return []int{}
}

func main() {
	var nums = []int{2, 7, 11, 15}
	// 1. Two sum
	var target int = 9
	TwoSum(nums, target)

}
