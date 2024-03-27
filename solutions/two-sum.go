package solutions

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
