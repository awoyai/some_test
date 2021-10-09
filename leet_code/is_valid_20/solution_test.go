package is_valid_20

import (
	"fmt"
	"testing"
)

func BenchmarkIsValid(b *testing.B) {
	isValid := IsValid("()[]{}")
	fmt.Printf("result: %v", isValid)
}