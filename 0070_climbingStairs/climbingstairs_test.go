package climbingstairs

import "testing"

func TestClimbingStairs(t *testing.T) {
	if res := climbStairs(3); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
