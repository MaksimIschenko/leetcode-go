package main

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

func main() {

	// 283. Move zeroes
	var nums = []int{2, 7, 11, 15}
	MoveZeroes(nums)

}
