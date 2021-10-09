package _0_is_valid

import (
	"fmt"
	"testing"
)

func BenchmarkIsValid(b *testing.B) {
	isValid := IsValid("()[]{}")
	fmt.Printf("result: %v", isValid)
}