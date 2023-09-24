package longestsubarrayof1safterdeletingoneelement

import "vudangkhoa2906-leetcode-go/common"

func longestSubarray(nums []int) int {
	var res, lenSubarray, lenPrevSubarray, tempRes int
	lenNums := len(nums)
	for idx := range nums {
		if nums[idx] == 0 {
			if idx > 0 {
				if nums[idx-1] == 0 {
					lenPrevSubarray = 0
				} else {
					if tempRes = lenSubarray + lenPrevSubarray; tempRes > res {
						res = tempRes
					}
					lenPrevSubarray = lenSubarray
				}
				lenSubarray = 0
			}
		} else {
			lenSubarray++
		}
	}
	if nums[lenNums-1] == 1 {
		if tempRes = lenSubarray + lenPrevSubarray; tempRes > res {
			res = tempRes
		}
	}
	return common.Max(res, len(nums)-1)
}
