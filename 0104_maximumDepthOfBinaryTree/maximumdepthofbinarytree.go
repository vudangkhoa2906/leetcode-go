package maximumdepthofbinarytree

import "vudangkhoa2906-leetcode-go/common"

func maxDepth(root *common.TreeNode) int {
	var maxDepthRec func(node *common.TreeNode) int
	maxDepthRec = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + common.Max(maxDepthRec(node.Left), maxDepthRec(node.Right))
	}
	return maxDepthRec(root)
}
