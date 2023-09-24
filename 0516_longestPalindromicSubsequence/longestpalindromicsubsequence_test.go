package longestpalindromicsubsequence

import "testing"

func TestLongestPalindromicSubsequence(t *testing.T) {
	s := "bbbab"
	if res := longestPalindromeSubseq(s); res != 4 {
		t.Errorf("Output: %d: Expected: %d\n", res, 4)
	}
}
