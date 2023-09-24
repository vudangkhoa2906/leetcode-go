package binarytreepreordertraversal

import "vudangkhoa2906-leetcode-go/common"

func preorderTraversal(root *common.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var top *common.TreeNode
	stack := []*common.TreeNode{root}
	for count := len(stack); count > 0; count = len(stack) {
		top, stack = stack[count-1], stack[:count-1]
		res = append(res, top.Val)
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}
