package _22_max_profit

// 动态规划
func maxProfit(prices []int) int {
	d0, d1 := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		d0, d1 = max(d0, d1+prices[i]), max(d1, d0-prices[i])
	}
	return d0
}

// tanxin
func maxProfitTx(prices []int) int {
	var res = 0
	for i := 1; i < len(prices); i++ {
		res += max(0, prices[i] - prices[i-1])
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
