package reverselinkedlist

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestReverseLinkedList(t *testing.T) {
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
	if res := reverseList(head); res.Val != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res.Val, 5)
	}
}
