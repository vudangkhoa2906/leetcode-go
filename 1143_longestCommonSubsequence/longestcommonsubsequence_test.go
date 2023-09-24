package longestcommonsubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	if res := longestCommonSubsequence(text1, text2); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
