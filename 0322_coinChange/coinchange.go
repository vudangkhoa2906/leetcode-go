package coinchange

import "math"

func coinChange(coins []int, amount int) int {
	res := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		res[a] = -1
	}
	var tempPrev int
	for a := 1; a <= amount; a++ {
		tempPrev = math.MaxInt
		for _, c := range coins {
			if c <= a && res[a-c] > -1 && tempPrev > res[a-c] {
				tempPrev = res[a-c]
			}
		}
		if tempPrev == math.MaxInt {
			res[a] = -1
		} else {
			res[a] = tempPrev + 1
		}
	}
	return res[amount]
}
