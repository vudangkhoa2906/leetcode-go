package wordsearch

import "testing"

func TestWordSearch(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	if res := exist(board, word); !res {
		t.Errorf("Output: %t: Expected: %t\n", res, true)
	}
}
