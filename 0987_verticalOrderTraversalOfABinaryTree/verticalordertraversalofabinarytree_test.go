package verticalordertraversalofabinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestVerticalOrderTraversalOfABinaryTree(t *testing.T) {
	root := &common.TreeNode{
		3,
		&common.TreeNode{9, nil, nil},
		&common.TreeNode{
			20,
			&common.TreeNode{15, nil, nil},
			&common.TreeNode{7, nil, nil},
		},
	}
	res := verticalTraversal(root)
	t.Logf("Output: %v\n", res)
}
