package partitionlist

import "vudangkhoa2906-leetcode-go/common"

func partition(head *common.ListNode, x int) *common.ListNode {
	if head == nil {
		return nil
	}
	curNode := head
	var part1Head, part2Head, part1CurNode, part2CurNode, nextNode *common.ListNode
	for curNode != nil {
		nextNode = curNode.Next
		if curNode.Val < x {
			if part1Head == nil {
				part1Head, part1CurNode = curNode, curNode
			} else {
				part1CurNode.Next, part1CurNode = curNode, curNode
			}
		} else {
			if part2Head == nil {
				part2Head, part2CurNode = curNode, curNode
			} else {
				part2CurNode.Next, part2CurNode = curNode, curNode
			}
		}
		curNode = nextNode
	}
	if part1Head == nil {
		return part2Head
	} else {
		part1CurNode.Next = part2Head
		if part2Head != nil {
			part2CurNode.Next = nil
		}
		return part1Head
	}
}
