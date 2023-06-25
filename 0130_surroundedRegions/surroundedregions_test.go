package surroundedregions

import (
	"fmt"
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
	fmt.Println(board)
}
