package binarytreezigzaglevelordertraversal

import "vudangkhoa2906-leetcode-go/common"

func zigzagLevelOrder(root *common.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*common.TreeNode{root}
	isOddLevel := true
	var tempRes []int
	for count := len(queue); count > 0; count = len(queue) {
		tempRes = make([]int, count)
		for c := 0; c < count; c++ {
			tempRes[c] = queue[c].Val
			if isOddLevel {
				if queue[count-1-c].Right != nil {
					queue = append(queue, queue[count-1-c].Right)
				}
				if queue[count-1-c].Left != nil {
					queue = append(queue, queue[count-1-c].Left)
				}
			} else {
				if queue[count-1-c].Left != nil {
					queue = append(queue, queue[count-1-c].Left)
				}
				if queue[count-1-c].Right != nil {
					queue = append(queue, queue[count-1-c].Right)
				}
			}
		}
		queue = queue[count:]
		res = append(res, tempRes)
		isOddLevel = !isOddLevel
	}
	return res
}
