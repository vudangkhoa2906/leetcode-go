package binarytreepostordertraversal

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBinaryTreePostorderTraversal(t *testing.T) {
	root := &common.TreeNode{
		1,
		nil,
		&common.TreeNode{
			2,
			&common.TreeNode{3, nil, nil},
			nil,
		},
	}
	res := postorderTraversal(root)
	t.Logf("Output: %v\n", res)
}
