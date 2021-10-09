package _11_min_depth


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MinDepth 递归
// info
//	解答成功:
//	执行耗时:176 ms,击败了91.36% 的Go用户
//	内存消耗:18.5 MB,击败了58.03% 的Go用户
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := MinDepth(root.Left)
	rightHeight := MinDepth(root.Right)
	return 1 + min(leftHeight, rightHeight)
}

func min(x int, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x > y {
		return y
	}
	return x
}
