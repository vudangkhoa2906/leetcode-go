package jumpgameiii

import "testing"

func TestJumpGameIII(t *testing.T) {
	arr := []int{0, 3, 0, 6, 3, 3, 4}
	start := 6
	if res := canReach(arr, start); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
