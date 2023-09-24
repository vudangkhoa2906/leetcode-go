package cherrypickupii

import "testing"

func TestCherryPickupII(t *testing.T) {
	grid := [][]int{
		{1, 0, 0, 0, 0, 0, 1},
		{2, 0, 0, 0, 0, 3, 0},
		{2, 0, 9, 0, 0, 0, 0},
		{0, 3, 0, 5, 4, 0, 0},
		{1, 0, 2, 3, 0, 0, 6},
	}
	if res := cherryPickup(grid); res != 28 {
		t.Errorf("Output: %d: Expected: %d\n", res, 28)
	}
}
