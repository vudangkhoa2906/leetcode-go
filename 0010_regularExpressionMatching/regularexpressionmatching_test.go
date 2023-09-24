package regularexpressionmatching

import (
	"testing"
)

func TestRegularExpressionMatching(t *testing.T) {
	s := "ab"
	p := ".*"
	if res := isMatch(s, p); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
