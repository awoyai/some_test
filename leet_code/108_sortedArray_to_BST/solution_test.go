package _08_sortedArray_to_BST

import (
	"fmt"
	"testing"
)

func BenchmarkSortedArrayToBST(b *testing.B) {
	node := SortedArrayToBST([]int{1, 2, 3, 4})
	fmt.Printf("%+v", node)
}