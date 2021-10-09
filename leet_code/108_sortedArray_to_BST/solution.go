package _08_sortedArray_to_BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  helper(nums, left, mid - 1),
		Right: helper(nums, mid +1, right),
	}
}

// 解法缺陷： 重新切片会占用资源和时间
//func SortedArrayToBST(nums []int) *TreeNode {
//	if len(nums) == 0 {
//		return nil
//	}
//	if len(nums) == 1 {
//		return &TreeNode{
//			Val: nums[0],
//		}
//	}
//	fmt.Println(len(nums))
//	mid := len(nums) / 2
//	return &TreeNode{
//		Val:   nums[mid],
//		Left:  SortedArrayToBST(nums[:mid]),
//		Right: SortedArrayToBST(nums[mid+1:]),
//	}
//}