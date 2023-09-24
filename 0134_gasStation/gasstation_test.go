package gasstation

import "testing"

func TestGasStation(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	if res := canCompleteCircuit(gas, cost); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
