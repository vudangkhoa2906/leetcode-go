package permutations

import "testing"

func TestPermutations(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	t.Logf("Output: %v\n", res)
}
