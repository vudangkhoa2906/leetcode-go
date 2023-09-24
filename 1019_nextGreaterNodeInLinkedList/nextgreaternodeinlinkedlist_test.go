package nextgreaternodeinlinkedlist

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestNextGreaterNodeInLinkedList(t *testing.T) {
	head := &common.ListNode{
		2,
		&common.ListNode{
			1,
			&common.ListNode{5, nil},
		},
	}
	res := nextLargerNodes(head)
	t.Logf("Output: %v\n", res)
}
