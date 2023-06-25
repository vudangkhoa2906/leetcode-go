package binarytreemaximumpathsum

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestBinaryTreeMaximumPathSum(t *testing.T) {
	root := &common.TreeNode{
		Val:  -10,
		Left: &common.TreeNode{Val: 9, Left: nil, Right: nil},
		Right: &common.TreeNode{
			Val:   20,
			Left:  &common.TreeNode{Val: 15, Left: nil, Right: nil},
			Right: &common.TreeNode{Val: 7, Left: nil, Right: nil}},
	}
	if res := maxPathSum(root); res != 42 {
		t.Errorf("Output: %d: Expected: 42\n", res)
	}
}
