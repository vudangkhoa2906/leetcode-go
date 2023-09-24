package lowestcommonancestorofabinarytree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestLowestCommonAncestorOfABinaryTree(t *testing.T) {
	root := &common.TreeNode{
		3,
		&common.TreeNode{
			5,
			&common.TreeNode{6, nil, nil},
			&common.TreeNode{
				2,
				&common.TreeNode{7, nil, nil},
				&common.TreeNode{4, nil, nil},
			},
		},
		&common.TreeNode{
			1,
			&common.TreeNode{0, nil, nil},
			&common.TreeNode{8, nil, nil},
		},
	}
	p, q := root.Left, root.Left.Right.Right
	res := lowestCommonAncestor(root, p, q)
	t.Logf("Output: %v\n", res)
}
