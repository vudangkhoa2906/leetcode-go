package jumpgameii

import "testing"

func TestJumpGame(t *testing.T) {
	nums := []int{2, 3, 1, 1, 3}
	if res := jump(nums); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
