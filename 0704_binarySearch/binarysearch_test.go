package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	if res := search(nums, target); res != 4 {
		t.Errorf("Output: %d: Expected: %d\n", res, 4)
	}
}
