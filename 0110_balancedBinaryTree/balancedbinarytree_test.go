package balancedbinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBalancedBinaryTree(t *testing.T) {
	root := &common.TreeNode{
		3,
		&common.TreeNode{9, nil, nil},
		&common.TreeNode{
			20,
			&common.TreeNode{15, nil, nil},
			&common.TreeNode{7, nil, nil},
		},
	}
	if res := isBalanced(root); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
