package binarytreerightsideview

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBinaryTreeRightSideView(t *testing.T) {
	root := &common.TreeNode{
		1,
		&common.TreeNode{
			2,
			nil,
			&common.TreeNode{5, nil, nil},
		},
		&common.TreeNode{
			3,
			nil,
			&common.TreeNode{4, nil, nil},
		},
	}
	res := rightSideView(root)
	t.Logf("Output: %v\n", res)
}
