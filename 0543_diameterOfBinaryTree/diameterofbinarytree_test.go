package diameterofbinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &common.TreeNode{
		1,
		&common.TreeNode{
			2,
			&common.TreeNode{4, nil, nil},
			&common.TreeNode{5, nil, nil},
		},
		&common.TreeNode{3, nil, nil},
	}
	if res := diameterOfBinaryTree(root); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
