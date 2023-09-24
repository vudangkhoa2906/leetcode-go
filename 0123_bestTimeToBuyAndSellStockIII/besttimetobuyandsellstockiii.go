package besttimetobuyandsellstockiii

import "vudangkhoa2906-leetcode-go/common"

func maxProfit(prices []int) int {
	len := len(prices)
	dp := make([][][]int, len+1)
	for idx := 0; idx <= len; idx++ {
		dp[idx] = make([][]int, 2)
		for canBuy := 0; canBuy < 2; canBuy++ {
			dp[idx][canBuy] = make([]int, 3)
		}
	}

	var res int
	for idx := len - 1; idx >= 0; idx-- {
		for canBuy := 0; canBuy < 2; canBuy++ {
			for cap := 2; cap > 0; cap-- {
				if canBuy == 0 {
					res = common.Max(
						prices[idx]+dp[idx+1][1][cap-1], /* sell */
						dp[idx+1][0][cap],               /* not sell */
					)
				} else {
					res = common.Max(
						-prices[idx]+dp[idx+1][0][cap], /* buy */
						dp[idx+1][1][cap],              /* not buy */
					)
				}
				dp[idx][canBuy][cap] = res
			}
		}
	}
	return dp[0][1][2]
}
