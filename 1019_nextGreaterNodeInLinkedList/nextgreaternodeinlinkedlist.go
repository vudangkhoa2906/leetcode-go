package nextgreaternodeinlinkedlist

import "vudangkhoa2906-leetcode-go/common"

func nextLargerNodes(head *common.ListNode) []int {
	curNode := common.ReverseLinkedList(head)
	size := common.LenLinkedList(curNode)
	res := make([]int, size)
	var stack []int
	curIdx := size - 1
	for curNode != nil {
		for count := len(stack); count > 0 && stack[count-1] <= curNode.Val; count = len(stack) {
			stack = stack[:count-1]
		}
		if count := len(stack); count > 0 {
			res[curIdx] = stack[count-1]
		} else {
			res[curIdx] = 0
		}
		stack = append(stack, curNode.Val)
		curIdx--
		curNode = curNode.Next
	}
	return res
}
