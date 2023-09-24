package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var res, resInc int
	len := len(prices)
	if len <= 1 {
		return res
	}
	min := prices[0]
	for idx := 1; idx < len; idx++ {
		resInc = prices[idx] - min
		if res < resInc {
			res = resInc
		}
		if min > prices[idx] {
			min = prices[idx]
		}
	}
	return res
}
