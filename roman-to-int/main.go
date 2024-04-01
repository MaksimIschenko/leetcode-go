package main

import (
	"fmt"
)

func RomanToInt(s string) int {

	var result int
	var romanMap = map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000,
	}

	var lenght int = len(s) - 1
	var current int
	var previos int

	for j := range s {
		current = romanMap[string(s[lenght-j])]
		if j == 0 {
			result += current
		} else {
			if current < previos {
				result -= current
			} else {
				result += current
			}
		}
		previos = current
	}
	return result
}

func main() {
	// 13. Roman to Integer
	var romanNumber string = "MCMXCIV"
	result := RomanToInt(romanNumber)
	fmt.Println(result)
}
