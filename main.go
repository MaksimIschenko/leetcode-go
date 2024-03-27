package main

import (
	"fmt"
	"leetcode-go/solutions"
)

func main() {

	// Move zeroes
	var nums = []int{2, 7, 11, 15}
	solutions.MoveZeroes(nums)

	// 1. Two sum
	var target int = 9
	solutions.TwoSum(nums, target)

	// 13. Roman to Integer
	var romanNumber string = "MCMXCIV"
	result := solutions.RomanToInt(romanNumber)
	fmt.Println(result)

}
