package maximumabsolutesumofanysubarray

import "vudangkhoa2906-leetcode-go/common"

func maxAbsoluteSum(nums []int) int {
	res, maxInc, minInc := common.Abs(nums[0]), nums[0], nums[0]
	for idx := 1; idx < len(nums); idx++ {
		if maxInc > 0 {
			maxInc += nums[idx]
		} else {
			maxInc = nums[idx]
		}
		if minInc < 0 {
			minInc += nums[idx]
		} else {
			minInc = nums[idx]
		}
		if maxInc > 0 && res < maxInc {
			res = maxInc
		}
		if minInc < 0 && res < -minInc {
			res = -minInc
		}
	}
	return res
}
