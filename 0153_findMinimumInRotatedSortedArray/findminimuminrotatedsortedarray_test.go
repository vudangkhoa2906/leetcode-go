package findminimuminrotatedsortedarray

import "testing"

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	if res := findMin(nums); res != 0 {
		t.Errorf("Output: %d: Expected: %d\n", res, 0)
	}
}
