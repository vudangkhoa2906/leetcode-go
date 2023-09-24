package besttimetobuyandsellstockwithcooldown

import "testing"

func TestBestTimeToBuyAndSellStockWithCooldown(t *testing.T) {
	prices := []int{1, 2, 3, 0, 2}
	if res := maxProfit(prices); res != 3 {
		t.Errorf("Output: %d: Expected: %d\n", res, 3)
	}
}
