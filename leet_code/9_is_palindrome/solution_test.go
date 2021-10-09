package __is_palindrome

import (
	"fmt"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	isPalindrome := IsPalindrome(1212)
	if !isPalindrome {
		fmt.Printf("result: %v", isPalindrome)
	}
}
