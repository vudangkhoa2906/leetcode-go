package removeduplicatesfromsortedarray

import "testing"

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	if res := removeDuplicates(nums); res != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
	t.Logf("Output: %v\n", nums)
}
