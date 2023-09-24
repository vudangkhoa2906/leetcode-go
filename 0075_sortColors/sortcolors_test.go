package sortcolors

import "testing"

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	t.Logf("Output: %v\n", nums)
}
