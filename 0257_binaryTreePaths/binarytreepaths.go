package binarytreepaths

import (
	"strconv"
	"strings"
	"vudangkhoa2906-leetcode-go/common"
)

func binaryTreePaths(root *common.TreeNode) []string {
	var res []string
	var binaryTreePathsRec func(node *common.TreeNode) []strings.Builder
	binaryTreePathsRec = func(node *common.TreeNode) []strings.Builder {
		var tempRes []strings.Builder
		if node == nil {
			return tempRes
		}
		if node.Left == nil && node.Right == nil {
			var tempSb strings.Builder
			tempSb.WriteString(strconv.Itoa(node.Val))
			tempRes = append(tempRes, tempSb)
			return tempRes
		}
		for _, sb := range binaryTreePathsRec(node.Left) {
			var tempSb strings.Builder
			tempSb.WriteString(strconv.Itoa(node.Val))
			tempSb.WriteString("->")
			tempSb.WriteString(sb.String())
			tempRes = append(tempRes, tempSb)
		}
		for _, sb := range binaryTreePathsRec(node.Right) {
			var tempSb strings.Builder
			tempSb.WriteString(strconv.Itoa(node.Val))
			tempSb.WriteString("->")
			tempSb.WriteString(sb.String())
			tempRes = append(tempRes, tempSb)
		}
		return tempRes
	}
	for _, tempSb := range binaryTreePathsRec(root) {
		res = append(res, tempSb.String())
	}
	return res
}
