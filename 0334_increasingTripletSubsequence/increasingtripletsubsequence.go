package increasingtripletsubsequence

import "vudangkhoa2906-leetcode-go/common"

func increasingTriplet(nums []int) bool {
	lenNums := len(nums)
	res := []int{nums[0]}
	for idx := 1; idx < lenNums; idx++ {
		if res[len(res)-1] < nums[idx] {
			res = append(res, nums[idx])
			if len(res) == 3 {
				return true
			}
		} else {
			lowerBound := common.LowerBound(res, nums[idx])
			res[lowerBound] = nums[idx]
		}
	}
	return false
}
