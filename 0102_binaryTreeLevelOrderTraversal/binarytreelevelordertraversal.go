package binarytreelevelordertraversal

import "vudangkhoa2906-leetcode-go/common"

func levelOrder(root *common.TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	queue := []*common.TreeNode{root}
	var tempRes []int

	for levelCount := len(queue); levelCount > 0; levelCount = len(queue) {
		tempRes = make([]int, levelCount)
		for c := 0; c < levelCount; c++ {
			tempRes[c] = queue[c].Val
			if queue[c].Left != nil {
				queue = append(queue, queue[c].Left)
			}
			if queue[c].Right != nil {
				queue = append(queue, queue[c].Right)
			}
		}
		res = append(res, tempRes)
		queue = queue[levelCount:]
	}
	return res
}
