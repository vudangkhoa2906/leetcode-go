package besttimetobuyandsellstockiv

import "testing"

func TestBestTimeToBuyAndSellStockIV(t *testing.T) {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	if res := maxProfit(k, prices); res != 7 {
		t.Errorf("Output: %d: Expected: %d\n", res, 7)
	}
}
