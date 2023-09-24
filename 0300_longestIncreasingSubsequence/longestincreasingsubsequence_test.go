package longestincreasingsubsequence

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	if res := lengthOfLIS(nums); res != 4 {
		t.Errorf("Output: %d: Expected: %d\n", res, 4)
	}
}
