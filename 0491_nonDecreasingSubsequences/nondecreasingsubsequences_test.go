package nondecreasingsubsequences

import "testing"

func TestNonDecreasingSubsequences(t *testing.T) {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	t.Logf("Output: %v\n", res)
}
