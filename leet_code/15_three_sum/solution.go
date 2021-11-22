package _5_three_sum

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	thrid := n-1
	for first := 0; first < n; first++ {
		if first > 0 && nums[first - 1] == nums[first] {
			continue
		}
		target := -nums[first]
		for second := first + 1; second < n; second++ {
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			for thrid > second && nums[second] + nums[thrid] > target {
				thrid--
			}
			if thrid == second {
				break
			}
			if nums[first] + nums[second] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[thrid]})
			}
		}
	}
	return ans
}