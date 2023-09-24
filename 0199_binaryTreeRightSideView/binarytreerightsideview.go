package binarytreerightsideview

import "vudangkhoa2906-leetcode-go/common"

func rightSideView(root *common.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	queue := []*common.TreeNode{root}
	for count := len(queue); count > 0; count = len(queue) {
		for c := 0; c < count; c++ {
			if queue[c].Left != nil {
				queue = append(queue, queue[c].Left)
			}
			if queue[c].Right != nil {
				queue = append(queue, queue[c].Right)
			}
		}
		res = append(res, queue[count-1].Val)
		queue = queue[count:]
	}
	return res
}
