package pathsumii

import "vudangkhoa2906-leetcode-go/common"

func pathSum(root *common.TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var bt []int
	var pathSumRec func(node *common.TreeNode, target int)
	pathSumRec = func(node *common.TreeNode, target int) {
		if node.Left == nil && node.Right == nil {
			if node.Val == target {
				tempRes := make([]int, 0, len(bt)+1)
				tempRes = append(tempRes, bt...)
				tempRes = append(tempRes, node.Val)
				res = append(res, tempRes)
			}
		} else {
			bt = append(bt, node.Val)
			if node.Left != nil {
				pathSumRec(node.Left, target-node.Val)
			}
			if node.Right != nil {
				pathSumRec(node.Right, target-node.Val)
			}
			bt = bt[:len(bt)-1]
		}
	}
	pathSumRec(root, targetSum)
	return res
}
