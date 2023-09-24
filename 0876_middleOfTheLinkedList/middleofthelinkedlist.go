package middleofthelinkedlist

import "vudangkhoa2906-leetcode-go/common"

func middleNode(head *common.ListNode) *common.ListNode {
	var size int
	curNode := head
	for curNode != nil {
		size++
		curNode = curNode.Next
	}
	size = size / 2
	curNode = head
	for count := 0; count < size; count++ {
		curNode = curNode.Next
	}
	return curNode
}
