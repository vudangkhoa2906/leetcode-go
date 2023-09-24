package outofboundarypaths

import "testing"

func TestOutOfBoundaryPaths(t *testing.T) {
	m, n, maxMove, startRow, startColumn := 1, 3, 3, 0, 1
	if res := findPaths(m, n, maxMove, startRow, startColumn); res != 12 {
		t.Errorf("Output: %d: Expected: %d\n", res, 12)
	}
}
