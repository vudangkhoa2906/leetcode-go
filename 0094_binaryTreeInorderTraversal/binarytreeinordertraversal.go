package binarytreeinordertraversal

import "vudangkhoa2906-leetcode-go/common"

func inorderTraversal(root *common.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	node := root
	var count int
	var stack []*common.TreeNode
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			if len(stack) == 0 {
				break
			}
			count = len(stack)
			node, stack = stack[count-1], stack[:count-1]
			res = append(res, node.Val)
			node = node.Right
		}
	}
	return res
}
