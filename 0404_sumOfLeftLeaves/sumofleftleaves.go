package sumofleftleaves

import "vudangkhoa2906-leetcode-go/common"

func sumOfLeftLeaves(root *common.TreeNode) int {
	var sumOfLeftLeavesRec func(node *common.TreeNode, isLeftNode bool) int
	sumOfLeftLeavesRec = func(node *common.TreeNode, isLeftNode bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if isLeftNode == true {
				return node.Val
			} else {
				return 0
			}
		}
		return sumOfLeftLeavesRec(node.Left, true) + sumOfLeftLeavesRec(node.Right, false)
	}
	return sumOfLeftLeavesRec(root, false)
}
