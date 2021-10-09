package _12_has_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// info
//			解答成功:
//			执行耗时:0 ms,击败了100.00% 的Go用户
//			内存消耗:4.6 MB,击败了62.71% 的Go用户

func HasPathSum(root *TreeNode, targetSum int) bool {
	return helper(root, targetSum, 0)
}

func helper(root *TreeNode, targetSum int, sum int) bool {
	if root == nil {
		return false
	}
	sum+=root.Val
	if root.Left == nil && root.Right == nil {
		return sum == targetSum
	}
	var b bool = false
	if root.Left != nil {
		b = helper(root.Left, targetSum, sum) || b
	}
	if root.Right != nil {
		b = helper(root.Right, targetSum, sum) || b
	}
	return b
}
