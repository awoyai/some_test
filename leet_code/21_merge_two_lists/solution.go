package mergeTwoLists

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start

type ListNode struct {
	Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	val := 0
	if list1.Val <= list2.Val {
		val = list1.Val
		list1 = list1.Next
	} else {
		val = list2.Val
		list2 = list2.Next
	}
	return &ListNode{Val: val, Next: mergeTwoLists(list1, list2)}
}
// @lc code=end

