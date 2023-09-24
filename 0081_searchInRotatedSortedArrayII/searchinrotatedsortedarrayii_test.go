package searchinrotatedsortedarrayii

import "testing"

func TestSearchInRotatedSortedArrayII(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1}
	target := 2
	if res := search(nums, target); res != true {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
