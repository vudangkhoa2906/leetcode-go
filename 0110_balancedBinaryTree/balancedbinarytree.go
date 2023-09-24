package balancedbinarytree

import (
	"vudangkhoa2906-leetcode-go/common"
)

func isBalanced(root *common.TreeNode) bool {
	var isBalancedRec func(node *common.TreeNode) (int, bool)
	isBalancedRec = func(node *common.TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		maxDepthLeft, isBalancedLeft := isBalancedRec(node.Left)
		if !isBalancedLeft {
			return 0, false
		}
		maxDepthRight, isBalancedRight := isBalancedRec(node.Right)
		if !isBalancedRight || common.Abs(maxDepthLeft-maxDepthRight) > 1 {
			return 0, false
		}
		return 1 + common.Max(maxDepthLeft, maxDepthRight), true
	}
	_, res := isBalancedRec(root)
	return res
}
