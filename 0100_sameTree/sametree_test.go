package sametree

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestSameTree(t *testing.T) {
	p := &common.TreeNode{
		1,
		&common.TreeNode{2, nil, nil},
		nil,
	}
	q := &common.TreeNode{
		1,
		nil,
		&common.TreeNode{2, nil, nil},
	}
	if res := isSameTree(p, q); res {
		t.Errorf("Output: %t: Expected: %t\n", res, false)
	}
}
