package _10_is_balanced

import (
	"fmt"
	"testing"
)

func BenchmarkIsBalanced(b *testing.B) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   2,
				Left:  nil,
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	isBalanced := IsBalanced(root)
	fmt.Printf("result: %v", isBalanced)
}