package countsquaresubmatriceswithallones

import "testing"

func TestCountSquareSubmatricesWithAllOnes(t *testing.T) {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	if res := countSquares(matrix); res != 15 {
		t.Errorf("Output: %d: Expected: %d\n", res, 15)
	}
}
