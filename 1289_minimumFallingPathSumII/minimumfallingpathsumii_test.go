package minimumfallingpathsumii

import "testing"

func TestMinimumFallingPathSumII(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if res := minFallingPathSum(matrix); res != 13 {
		t.Errorf("Output: %d: Expected: %d\n", res, 13)
	}
}
