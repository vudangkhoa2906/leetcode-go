package binarytreemaximumpathsum

import "vudangkhoa2906-leetcode-go/common"

func maxPathSum(root *common.TreeNode) int {
	var maxPathSumRec func(node *common.TreeNode) (int, int)
	maxPathSumRec = func(node *common.TreeNode) (int, int) {
		var max, maxInc int
		if node == nil {
			return max, maxInc
		}
		maxLeft, maxIncLeft := maxPathSumRec(node.Left)
		maxRight, maxIncRight := maxPathSumRec(node.Right)
		maxInc = node.Val + common.Max(common.Max(maxIncLeft, maxIncRight), 0)
		maxIncBoth := node.Val
		if maxIncLeft > 0 {
			maxIncBoth += maxIncLeft
		}
		if maxIncRight > 0 {
			maxIncBoth += maxIncRight
		}
		max = maxIncBoth
		if node.Left != nil && maxLeft > max {
			max = maxLeft
		}
		if node.Right != nil && maxRight > max {
			max = maxRight
		}
		return max, maxInc
	}
	res, _ := maxPathSumRec(root)
	return res
}
