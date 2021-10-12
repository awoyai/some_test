package _21_max_profit

import "math"

func MaxProfit(prices []int) int {
	var min int = math.MaxInt32
	var gap int = 0
	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i] - min > gap {
			gap = prices[i] - min
		}
	}
	return gap
}
