package maximumnumberofpointswithcost

import "testing"

func TestMaximumNumberOfPointsWithCost(t *testing.T) {
	points := [][]int{
		{0, 3, 0, 4, 2},
		{5, 4, 2, 4, 1},
		{5, 0, 0, 5, 1},
		{2, 0, 1, 0, 3},
	}
	if res := maxPoints(points); res != 15 {
		t.Errorf("Output: %d: Expected: %d\n", res, 15)
	}
}
