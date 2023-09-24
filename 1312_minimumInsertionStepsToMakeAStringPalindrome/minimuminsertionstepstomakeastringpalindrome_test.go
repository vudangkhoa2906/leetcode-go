package minimuminsertionstepstomakeastringpalindrome

import "testing"

func TestMinimumInsertionStepsToMakeAStringPalindrome(t *testing.T) {
	s := "mbadm"
	if res := minInsertions(s); res != 2 {
		t.Errorf("Output: %d: Expected: %d\n", res, 2)
	}
}
