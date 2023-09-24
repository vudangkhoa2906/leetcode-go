package longestincreasingsubsequence

import "vudangkhoa2906-leetcode-go/common"

func lengthOfLIS(nums []int) int {
	size := len(nums)
	res := []int{nums[0]}
	for idx := 1; idx < size; idx++ {
		if sizeRes := len(res); sizeRes == 0 || res[sizeRes-1] < nums[idx] {
			res = append(res, nums[idx])
		} else {
			tempIdx := common.LowerBound(res, nums[idx])
			res[tempIdx] = nums[idx]
		}
	}
	return len(res)
}
