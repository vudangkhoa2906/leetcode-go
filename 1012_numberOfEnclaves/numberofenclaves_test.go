package numberofenclaves

import (
	"testing"
)

func TestNumberOfEnclaves(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	if res := numEnclaves(grid); res != 3 {
		t.Errorf("Output: %d: Expected: 3\n", res)
	}
}
