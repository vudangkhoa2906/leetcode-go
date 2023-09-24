package maximumwidthofbinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestMaximumWidthOfBinaryTree(t *testing.T) {
	root := &common.TreeNode{
		1,
		&common.TreeNode{
			3,
			&common.TreeNode{
				5,
				&common.TreeNode{6, nil, nil},
				nil,
			},
			nil,
		},
		&common.TreeNode{
			2,
			nil,
			&common.TreeNode{
				9,
				&common.TreeNode{7, nil, nil},
				nil,
			},
		},
	}
	if res := widthOfBinaryTree(root); res != 7 {
		t.Errorf("Output: %d: Expected: %d\n", res, 7)
	}
}
