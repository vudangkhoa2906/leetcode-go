package cherrypickup

import "testing"

func TestCherryPickup(t *testing.T) {
	grid := [][]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, 1, 1},
	}
	if res := cherryPickup(grid); res != 5 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
}
