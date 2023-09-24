package combinationsumii

import "testing"

func TestCombinationSumII(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	t.Logf("Output: %v\n", combinationSum2(candidates, target))
}
