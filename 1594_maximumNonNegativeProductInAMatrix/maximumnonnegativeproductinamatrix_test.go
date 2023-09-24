package maximumnonnegativeproductinamatrix

import "testing"

func TestMaximumNonNegativeProductInAMatrix(t *testing.T) {
	grid := [][]int{
		{-1, -2, -3},
		{-2, -3, -3},
		{-3, -3, -2},
	}
	if res := maxProductPath(grid); res != -1 {
		t.Errorf("Output: %d: Expected: %d\n", res, -1)
	}
}
