package surroundedregions

import (
	"testing"
)

func TestSurroundedRegions(t *testing.T) {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	t.Logf("Output: %v\n", board)
}
