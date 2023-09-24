package houserobberiii

import "vudangkhoa2906-leetcode-go/common"

func rob(root *common.TreeNode) int {
	dp := make(map[*common.TreeNode][]int)

	var maxRec func(node *common.TreeNode) (int, int)
	maxRec = func(node *common.TreeNode) (int, int) {
		var resExc, res int
		if node == nil {
			return resExc, res
		}
		if tempRes, prs := dp[node]; prs {
			return tempRes[0], tempRes[1]
		}

		maxExcLeft, maxLeft := maxRec(node.Left)
		maxExcRight, maxRight := maxRec(node.Right)
		resExc = maxLeft + maxRight
		res = common.Max(
			resExc,
			node.Val+maxExcRight+maxExcLeft,
		)
		dp[node] = make([]int, 2)
		dp[node][0], dp[node][1] = resExc, res
		return resExc, res
	}

	_, res := maxRec(root)
	return res
}
