package verticalordertraversalofabinarytree

import (
	"sort"
	"vudangkhoa2906-leetcode-go/common"
)

func verticalTraversal(root *common.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	tempRes := make(map[common.Pos][]int)
	var enrichTempRes func(node *common.TreeNode, row, col int)
	enrichTempRes = func(node *common.TreeNode, row, col int) {
		if node == nil {
			return
		}
		tempRes[common.Pos{R: row, C: col}] = append(tempRes[common.Pos{R: row, C: col}], node.Val)
		enrichTempRes(node.Left, row+1, col-1)
		enrichTempRes(node.Right, row+1, col+1)
	}
	enrichTempRes(root, 0, 0)
	keys := make([]common.Pos, len(tempRes))
	var idx int
	for key := range tempRes {
		keys[idx] = key
		idx++
	}
	sort.Slice(keys, func(idx1, idx2 int) bool {
		return keys[idx1].C < keys[idx2].C || keys[idx1].C == keys[idx2].C && keys[idx1].R < keys[idx2].R
	})
	col := keys[0].C - 1
	idx = -1
	for _, key := range keys {
		if key.C > col {
			col = key.C
			res = append(res, nil)
			idx++
		}
		sort.Ints(tempRes[key])
		res[idx] = append(res[idx], tempRes[key]...)
	}
	return res
}
