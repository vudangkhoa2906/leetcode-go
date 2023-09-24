package detonatethemaximumbombs

import "testing"

func TestDetonateTheMaximumBombs(t *testing.T) {
	bombs := [][]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 4, 2},
		{4, 5, 3},
		{5, 6, 4},
	}
	if res := maximumDetonation(bombs); res != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
}
