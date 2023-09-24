package containerwithmostwater

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	if res := maxArea(height); res != 49 {
		t.Errorf("Output: %d: Expected: %d\n", res, 49)
	}
}
