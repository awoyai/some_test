package mergeTwoLists

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	node := mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	})
	fmt.Printf("result: %v", *node)
}