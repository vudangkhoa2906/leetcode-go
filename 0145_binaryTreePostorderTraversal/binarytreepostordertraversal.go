package binarytreepostordertraversal

import "vudangkhoa2906-leetcode-go/common"

func postorderTraversal(root *common.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack1 := []*common.TreeNode{root}
	var top *common.TreeNode
	var stack2 []*common.TreeNode

	for count := len(stack1); count > 0; count = len(stack1) {
		top, stack1 = stack1[count-1], stack1[:count-1]
		stack2 = append(stack2, top)
		if top.Left != nil {
			stack1 = append(stack1, top.Left)
		}
		if top.Right != nil {
			stack1 = append(stack1, top.Right)
		}
	}
	for count := len(stack2); count > 0; count = len(stack2) {
		top, stack2 = stack2[count-1], stack2[:count-1]
		res = append(res, top.Val)
	}
	return res
}
