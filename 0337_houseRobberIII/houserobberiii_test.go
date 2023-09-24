package houserobberiii

import (
	"testing"
	"vudangkhoa2906-leetcode-go/common"
)

func TestHouseRobberIII(t *testing.T) {
	root := &common.TreeNode{
		3,
		&common.TreeNode{
			2,
			nil,
			&common.TreeNode{3, nil, nil},
		},
		&common.TreeNode{
			3,
			nil,
			&common.TreeNode{1, nil, nil},
		},
	}
	if res := rob(root); res != 7 {
		t.Errorf("Output: %d: Expected: %d\n", res, 7)
	}
}
