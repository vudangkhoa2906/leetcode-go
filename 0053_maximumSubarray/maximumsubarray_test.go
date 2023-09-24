package maximumsubarray

import "testing"

func TestMaximumSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if res := maxSubArray(nums); res != 6 {
		t.Errorf("Output: %d: Expected: %v\n", res, 6)
	}
}
