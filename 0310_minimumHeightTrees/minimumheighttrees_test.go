package minimumheighttrees

import "testing"

func TestMinimumHeightTrees(t *testing.T) {
	n := 4
	edges := [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
	}
	t.Logf("Output: %v\n", findMinHeightTrees(n, edges))
}
