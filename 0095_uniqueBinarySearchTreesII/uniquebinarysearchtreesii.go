package uniquebinarysearchtreesii

import "vudangkhoa2906-leetcode-go/common"

func generateTrees(n int) []*common.TreeNode {
	var generateTreesRec func(lowVal, highVal int) []*common.TreeNode
	generateTreesRec = func(lowVal, highVal int) []*common.TreeNode {
		if lowVal > highVal {
			return []*common.TreeNode{nil}
		}
		var res []*common.TreeNode
		for val := lowVal; val <= highVal; val++ {
			for _, leftTree := range generateTreesRec(lowVal, val-1) {
				for _, rightTree := range generateTreesRec(val+1, highVal) {
					res = append(res, &common.TreeNode{
						val,
						leftTree,
						rightTree,
					})
				}
			}
		}
		return res
	}
	return generateTreesRec(1, n)
}
