package largestrectangleinhistogram

import "testing"

func TestLargestRectangleInHistogram(t *testing.T) {
	heights := []int{2, 2, 3, 2, 2}
	if res := largestRectangleArea(heights); res != 10 {
		t.Errorf("Output: %d: Expected: %d\n", res, 10)
	}
}
