package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	s := "abcd"
	wordDict := []string{"a", "abc", "b", "cd"}
	if res := wordBreak(s, wordDict); res != true {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
