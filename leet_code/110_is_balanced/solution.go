package _10_is_balanced

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	b, _ := helper(root, 0)
	return b
}

func helper(root *TreeNode, count int) (bool, int) {
	if root == nil {
		return true, count
	}
	count++
	leftBoolean, leftLen := helper(root.Left, 0)
	if !leftBoolean {
		return leftBoolean, 0
	}
	rightBoolean, rightLen := helper(root.Right, 0)
	if !rightBoolean {
		return rightBoolean, 0
	}
	if math.Abs(float64(leftLen - rightLen)) > 1{
		return false, 0
	}
	return true, count + max(leftLen, rightLen)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}