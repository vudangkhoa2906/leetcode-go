package wildcardmatching

import (
	"testing"
)

func TestWildcardMatching(t *testing.T) {
	s := "abcde"
	p := "*c?e**"
	if res := isMatch(s, p); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
