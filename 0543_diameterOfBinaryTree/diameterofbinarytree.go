package diameterofbinarytree

import (
	"vudangkhoa2906-leetcode-go/common"
)

func diameterOfBinaryTree(root *common.TreeNode) int {
	var diameterOfBinaryTreeRec func(node *common.TreeNode) (int, int)
	diameterOfBinaryTreeRec = func(node *common.TreeNode) (int, int) {
		var max, maxInc int
		if node == nil || (node.Left == nil && node.Right == nil) {
			return max, maxInc
		}
		maxLeft, maxIncLeft := diameterOfBinaryTreeRec(node.Left)
		maxRight, maxIncRight := diameterOfBinaryTreeRec(node.Right)
		maxInc = 1 + common.Max(maxIncLeft, maxIncRight)
		var maxIncBoth int
		if node.Right != nil {
			maxIncBoth += 1 + maxIncRight
		}
		if node.Left != nil {
			maxIncBoth += 1 + maxIncLeft
		}
		max = common.Max(common.Max(maxLeft, maxRight), maxIncBoth)
		return max, maxInc
	}
	res, _ := diameterOfBinaryTreeRec(root)
	return res
}
