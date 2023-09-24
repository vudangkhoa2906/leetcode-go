package partitionarrayintotwoarraystominimizesumdifference

import (
	"math"
)

func minimumDifference(nums []int) int {
	normalize(nums)
	lenNums := len(nums)
	res := math.MinInt
	var sumNums, sumBt int
	var lenBt int
	for _, val := range nums {
		sumNums += val
	}

	var minimumDifferenceRec func(idx int)
	minimumDifferenceRec = func(idx int) {
		if lenBt == lenNums/2 {
			if sumBt <= sumNums/2 && res < sumBt {
				res = sumBt
			}
		} else if idx <= lenNums/2+lenBt {
			lenBt++
			sumBt += nums[idx]
			minimumDifferenceRec(idx + 1)
			lenBt--
			sumBt -= nums[idx]
			minimumDifferenceRec(idx + 1)
		}
	}

	minimumDifferenceRec(0)
	return sumNums - 2*res
}

func normalize(nums []int) {
	minNums := math.MaxInt
	for _, val := range nums {
		if val < minNums {
			minNums = val
		}
	}
	if minNums < 0 {
		for idx := range nums {
			nums[idx] -= minNums
		}
	}
}
