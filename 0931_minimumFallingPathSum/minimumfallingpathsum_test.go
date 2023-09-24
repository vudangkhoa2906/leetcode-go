package minimumfallingpathsum

import "testing"

func TestMinimumFallingPathSum(t *testing.T) {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}
	if res := minFallingPathSum(matrix); res != 13 {
		t.Errorf("Output: %d: Expected: %d\n", res, 13)
	}
}
