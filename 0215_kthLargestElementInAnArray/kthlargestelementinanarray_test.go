package kthlargestelementinanarray

import "testing"

func TestKthLargestElementInAnArray(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	if res := findKthLargest(nums, k); res != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
}
