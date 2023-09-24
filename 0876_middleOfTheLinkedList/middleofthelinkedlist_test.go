package middleofthelinkedlist

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestMiddleOfTheLinkedList(t *testing.T) {
	head := &common.ListNode{
		1,
		&common.ListNode{
			2,
			&common.ListNode{
				3,
				&common.ListNode{
					4,
					&common.ListNode{5, nil},
				},
			},
		},
	}
	if res := middleNode(head); res.Val != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res.Val, 3)
	}
}
