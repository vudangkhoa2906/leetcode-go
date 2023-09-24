package laststoneweight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	if res := lastStoneWeight(stones); res != 1 {
		t.Errorf("Output: %d: Expected: %d\n", res, 1)
	}
}
