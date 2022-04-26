package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	next := head.Next
	var aray = []*ListNode{head}
	for next != nil {
		aray = append(aray, next)
		next = next.Next
	}
	len := len(aray)
	switch {
	case len == 1:
		return nil
	case len == n:
		return head.Next
	case n == 1 && len > 1:
		aray[len-2].Next = nil
	default:
		aray[len-n-1].Next = aray[len-n+1]
	}
	return head
}

// @lc code=end
