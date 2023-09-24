package permutationsii

import "testing"

func TestPermutationsII(t *testing.T) {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	t.Logf("Output: %v\n", res)
}
