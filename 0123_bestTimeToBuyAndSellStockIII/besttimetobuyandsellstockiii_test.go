package besttimetobuyandsellstockiii

import "testing"

func TestBestTimeToBuyAndSellStockIII(t *testing.T) {
	prices := []int{8, 3, 6, 2, 8, 8, 8, 4, 2, 0, 7, 2, 9, 4, 9}
	if res := maxProfit(prices); res != 15 {
		t.Errorf("Output: %d: Expected: %d\n", res, 15)
	}
}
