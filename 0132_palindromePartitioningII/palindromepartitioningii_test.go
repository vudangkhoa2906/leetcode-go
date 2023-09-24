package palindromepartitioningii

import "testing"

func TestPalindromePartitioningII(t *testing.T) {
	s := "aab"
	if res := minCut(s); res != 1 {
		t.Errorf("Output: %d: Expected: %d\n", res, 1)
	}
}
