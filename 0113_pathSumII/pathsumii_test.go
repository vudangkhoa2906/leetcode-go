package pathsumii

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestPathSumII(t *testing.T) {
	root := &common.TreeNode{
		5,
		&common.TreeNode{
			4,
			&common.TreeNode{
				11,
				&common.TreeNode{7, nil, nil},
				&common.TreeNode{2, nil, nil},
			},
			nil,
		},
		&common.TreeNode{
			8,
			&common.TreeNode{13, nil, nil},
			&common.TreeNode{
				4,
				&common.TreeNode{5, nil, nil},
				&common.TreeNode{1, nil, nil},
			},
		},
	}
	targetSum := 22
	res := pathSum(root, targetSum)
	t.Logf("Output: %v\n", res)
}
