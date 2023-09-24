package besttimetobuyandsellstock

import "testing"

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	prices := []int{1, 2}
	if res := maxProfit(prices); res != 1 {
		t.Errorf("Output: %d: Expected: %d\n", res, 5)
	}
}
