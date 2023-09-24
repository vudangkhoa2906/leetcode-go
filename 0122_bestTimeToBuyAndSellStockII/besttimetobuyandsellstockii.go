package besttimetobuyandsellstockii

func maxProfit(prices []int) int {
	len := len(prices)
	if len <= 1 {
		return 0
	}
	var res int
	for idx := 1; idx < len; idx++ {
		if prices[idx] > prices[idx-1] {
			res += prices[idx] - prices[idx-1]
		}
	}
	return res
}
