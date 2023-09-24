package maximumcandiesallocatedtokchildren

import "math"

func maximumCandies(candies []int, k int64) int {
	var sumCandies, tempChildren int64
	var midCandies int
	res, minCandies, maxCandies := math.MinInt, 1, math.MinInt
	for _, val := range candies {
		sumCandies += int64(val)
		if maxCandies < val {
			maxCandies = val
		}
	}
	if sumCandies < k {
		return 0
	}

	for minCandies <= maxCandies {
		midCandies, tempChildren = (minCandies+maxCandies)/2, 0
		for _, val := range candies {
			tempChildren += int64(val / midCandies)
		}
		if tempChildren >= k {
			if res < midCandies {
				res = midCandies
			}
			minCandies = midCandies + 1
		} else {
			maxCandies = midCandies - 1
		}
	}
	return res
}
