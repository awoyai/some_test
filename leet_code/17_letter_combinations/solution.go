package lettercombinations

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var (
	letter_map = []string{
		" ",    //0
        "",     //1
        "abc",  //2
        "def",  //3
        "ghi",  //4
        "jkl",  //5
        "mno",  //6
        "pqrs", //7
        "tuv",  //8
        "wxyz",  //9
	}
	res = []string{}
)

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	dpfs(digits, 0, "")
	return res
}

func dpfs(digits string, index int, s string) {
	if len(digits) <= index {
		res = append(res, s)
		return
	}
	letter := letter_map[digits[index] - '0']
	index++
	for i := range letter {
		dpfs(digits, index, s + string(letter[i]))
	}
}
// @lc code=end

