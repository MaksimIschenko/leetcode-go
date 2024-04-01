package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
