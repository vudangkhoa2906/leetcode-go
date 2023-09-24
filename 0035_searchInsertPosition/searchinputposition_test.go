package searchinputposition

import "testing"

func TestSearchInputPosition(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	if res := searchInsert(nums, target); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
