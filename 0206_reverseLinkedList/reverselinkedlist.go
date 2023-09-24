package reverselinkedlist

import "vudangkhoa2906-leetcode-go/common"

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	curNode := head
	var prevNode, nextNode *common.ListNode
	for curNode != nil {
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	return prevNode
}
