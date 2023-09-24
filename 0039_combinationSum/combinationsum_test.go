package combinationsum

import "testing"

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	t.Logf("Output: %v\n", combinationSum(candidates, target))
}
