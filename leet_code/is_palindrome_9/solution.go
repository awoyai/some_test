package is_palindrome_9

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	len := len(xStr)
	for i, _ := range xStr {
		if xStr[i] != xStr[len-1-i] {
			return false
		}
	}
	return true
}
