package jumpgame

import "testing"

func TestJumpGame(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	if res := canJump(nums); res {
		t.Errorf("Output: %t: Expected: %t\n", res, false)
	}
}
