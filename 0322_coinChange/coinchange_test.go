package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	if res := coinChange(coins, amount); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
