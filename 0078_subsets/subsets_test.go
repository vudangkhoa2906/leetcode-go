package subsets

import "testing"

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	t.Logf("Output: %v\n", res)
}
