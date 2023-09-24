package palindromepartitioning

import "testing"

func TestPalindromePartitioning(t *testing.T) {
	s := "aabb"
	res := partition(s)
	t.Logf("Output: %v\n", res)
}
