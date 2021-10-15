package __is_palindrome

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	len := len(xStr)
	for i := 0; i < len/2; i++ {
		if xStr[i] != xStr[len-1-i] {
			return false
		}
	}
	return true
}
