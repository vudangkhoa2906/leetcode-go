package searchinrotatedsortedarray

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	if res := search(nums, target); res != 4 {
		t.Errorf("Output: %d: Expected: %d\n", res, 4)
	}
}
