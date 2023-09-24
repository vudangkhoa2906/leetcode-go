package binarytreepreordertraversal

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBinaryTreePreorderTraversal(t *testing.T) {
	root := &common.TreeNode{
		1,
		nil,
		&common.TreeNode{
			2,
			&common.TreeNode{3, nil, nil},
			nil,
		},
	}
	res := preorderTraversal(root)
	t.Logf("Output: %v\n", res)
}
