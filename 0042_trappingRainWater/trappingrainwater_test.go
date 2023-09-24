package trappingrainwater

import "testing"

func TestTrappingRainWater(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	if res := trap(height); res != 6 {
		t.Errorf("Output: %d: Expected: %d\n", res, 6)
	}
}
