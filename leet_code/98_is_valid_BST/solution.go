package _8_is_valid_BST

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return lHelper(root.Left, root.Val, math.MinInt) && rHelper(root.Right, root.Val, math.MaxInt)
}

func lHelper(root *TreeNode, val, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= val || root.Val <= min {
		return false
	}
	return lHelper(root.Left, root.Val, min) && rHelper(root.Right, root.Val, val)
}

func rHelper(root *TreeNode, val, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= val || root.Val >= max{
		return false
	}
	return lHelper(root.Left, root.Val, val) && rHelper(root.Right, root.Val, max)
}


// isValidBSTPUB 官方解
func isValidBSTPUB(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}