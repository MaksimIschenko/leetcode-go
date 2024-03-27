package solutions

func MoveZeroes(nums []int) {
	var j int
	for i := range nums {
		j = i
		if nums[i] == 0 {
			for nums[j] == 0 && j < len(nums)-1 {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
