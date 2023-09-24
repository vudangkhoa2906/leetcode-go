package minimumpathcostinagrid

import "testing"

func TestMinimumPathCostInAGrid(t *testing.T) {
	grid := [][]int{{5, 3}, {4, 0}, {2, 1}}
	moveCost := [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}
	if res := minPathCost(grid, moveCost); res != 17 {
		t.Errorf("Output: %d: Expected: %d\n", res, 17)
	}
}
