package distinctsubsequences

import "testing"

func TestDistinctSubsequences(t *testing.T) {
	s := "babgbag"
	r := "bag"
	if res := numDistinct(s, r); res != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
}
