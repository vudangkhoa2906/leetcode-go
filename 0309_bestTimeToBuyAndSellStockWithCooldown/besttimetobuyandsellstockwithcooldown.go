package besttimetobuyandsellstockwithcooldown

import "vudangkhoa2906-leetcode-go/common"

func maxProfit(prices []int) int {
	len := len(prices)
	if len <= 1 {
		return 0
	}
	dp := make([][]int, len+2)
	for idx := 0; idx <= len+1; idx++ {
		dp[idx] = make([]int, 2)
	}
	var res int
	for idx := len - 1; idx >= 0; idx-- {
		for canBuy := 0; canBuy < 2; canBuy++ {
			if canBuy == 0 { /* cannot buy */
				res = common.Max(
					prices[idx]+dp[idx+2][1], /* sell */
					dp[idx+1][0],             /* not sell */
				)
			} else { /* can buy */
				res = common.Max(
					-prices[idx]+dp[idx+1][0], /* buy */
					dp[idx+1][1],              /* not buy */
				)
			}
			dp[idx][canBuy] = res
		}
	}
	return dp[0][1]
}
