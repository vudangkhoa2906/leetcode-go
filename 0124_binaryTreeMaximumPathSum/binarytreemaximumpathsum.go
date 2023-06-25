package binarytreemaximumpathsum

import "vudangkhoa2906-leetcode-go/common"

func maxPathSum(root *common.TreeNode) int {
	_, res := getMaxCur(root)
	return res
}

func getMaxCur(node *common.TreeNode) (int, int) {
	var maxIncCur, maxCur int
	if node == nil {
		return maxIncCur, maxCur
	}
	maxIncLeft, maxLeft := getMaxCur(node.Left)
	maxIncRight, maxRight := getMaxCur(node.Right)
	maxIncCur = node.Val + common.Max(0, common.Max(maxIncLeft, maxIncRight))
	maxCur = node.Val
	if maxIncLeft > 0 {
		maxCur += maxIncLeft
	}
	if maxIncRight > 0 {
		maxCur += maxIncRight
	}
	if node.Left != nil {
		maxCur = common.Max(maxCur, maxLeft)
	}
	if node.Right != nil {
		maxCur = common.Max(maxCur, maxRight)
	}
	return maxIncCur, maxCur
}
