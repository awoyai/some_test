package _66_plus_one

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start

func plusOne(digits []int) []int {
	var idx = len(digits) - 1
	for {
		if digits[idx] != 9 {
			digits[idx]++
			break
		}
		digits[idx] = 0
		if idx == 0 {
			digits = append([]int{1}, digits...)
			break
		}
		idx--
	}
	return digits
}

// @lc code=end
