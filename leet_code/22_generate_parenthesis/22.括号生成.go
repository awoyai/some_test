package main

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
var total = []string{}
func generateParenthesisDFS(n int) []string {
	if n == 0 {
		return nil
	}
	total = []string{}
	dfs(n, n, []byte{})
	return total
}

func dfs(left, right int, str []byte) {
	if left == 0 && right == 0 {
		total = append(total, string(str))
		return
	}
	if left > 0 {
		str = append(str, '(')
		left--
		dfs(left, right, str)
	} 
	if right > left {
		str = append(str, ')')
		right--
		dfs(left, right, str)
	}
}