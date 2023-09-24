package maximumdepthofbinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	root := &common.TreeNode{3,
		&common.TreeNode{9, nil, nil},
		&common.TreeNode{
			20,
			&common.TreeNode{15, nil, nil},
			&common.TreeNode{7, nil, nil},
		},
	}
	if res := maxDepth(root); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
