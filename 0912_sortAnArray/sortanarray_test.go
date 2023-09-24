package sortanarray

import "testing"

func TestSortAnArray(t *testing.T) {
	nums := []int{5, 1, 1, 2, 0, 0}
	sortArray(nums)
	t.Logf("Output: %v\n", nums)
}
