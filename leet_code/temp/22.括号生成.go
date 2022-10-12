package temp

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var total = [][]string{}
	total = append(total, []string{""}, []string{"()"})
	for i := 2; i <= n+1; i++ {
		l := []string{}
		for j := 0; j < i; j++ {
			now1 := total[j]
			now2 := total[i-j-1]
			for _, v1 := range now1 {
				for _, v2 := range now2 {
					el := "(" + v1 + ")" + v2
					l = append(l, el)
				}
			}
		}
		total = append(total, l)
	}
	return total[n]
}

// @lc code=end
