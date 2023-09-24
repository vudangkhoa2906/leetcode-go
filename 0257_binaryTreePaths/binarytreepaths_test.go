package binarytreepaths

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBinaryTreePaths(t *testing.T) {
	root := &common.TreeNode{
		1,
		&common.TreeNode{
			2,
			nil,
			&common.TreeNode{5, nil, nil},
		},
		&common.TreeNode{3, nil, nil},
	}
	res := binaryTreePaths(root)
	t.Logf("Output: %v\n", res)
}
