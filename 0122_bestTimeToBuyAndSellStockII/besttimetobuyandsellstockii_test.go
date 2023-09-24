package besttimetobuyandsellstockii

import "testing"

func TestBestTimeToBuyAndSellStockII(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	if res := maxProfit(prices); res != 7 {
		t.Errorf("Output: %d: Expected: %d\n", res, 7)
	}
}
