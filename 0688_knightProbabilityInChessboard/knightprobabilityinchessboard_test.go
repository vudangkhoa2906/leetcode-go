package knightprobabilityinchessboard

import "testing"

func TestKnightProbabilityInChessboard(t *testing.T) {
	n, k, row, column := 3, 2, 0, 0
	if res := knightProbability(n, k, row, column); res != 0.0625 {
		t.Errorf("Output: %f: Expected: %f\n", res, 0.0625)
	}
}
