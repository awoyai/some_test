package temp
/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res:= ""
	for i := 0; i < len(strs[0]); i++ {
		temp := strs[0][i]
		for j := 0; j < len(strs); j ++ {
			if len(strs[j]) <= i{
				return res
			}
			if temp != strs[j][i] {
				return res
			}
		}
		res += string(temp)
	}
	return res
}

// @lc code=end