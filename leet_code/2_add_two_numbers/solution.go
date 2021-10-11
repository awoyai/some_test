package __add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两者速度差不多 但迭代的内存更小一点
// 迭代
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var result *ListNode
	var count = 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + count
		sum, count = sum % 10, sum / 10
		if head == nil {
			head = &ListNode{
				Val:  sum,
			}
			result = head
		}else {
			result.Next = &ListNode{
				Val:  sum,
			}
			result = result.Next
		}
	}
	if count > 0 {
		result.Next = &ListNode{
			Val:  count,
			Next: nil,
		}
	}
	return head
}

// 递归
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

func helper(l1 *ListNode, l2 *ListNode, count int) *ListNode {
	if l1 == nil && l2 == nil && count == 0 {
		return nil
	}else if l1 == nil && l2 == nil && count != 0 {
		return &ListNode{
			Val:  count,
			Next: nil,
		}
	}
	var l1Next *ListNode
	var l2Next *ListNode
	val := 0
	switch {
	case l1 == nil:
		val = l2.Val + count
		l2Next = l2.Next
		l1Next = nil
	case l2 == nil:
		val = l1.Val + count
		l1Next = l1.Next
		l2Next = nil
	default:
		val = l1.Val + l2.Val + count
		l1Next = l1.Next
		l2Next = l2.Next
	}
	return &ListNode{
		Val:  val % 10,
		Next: helper(l1Next, l2Next, val / 10),
	}
}