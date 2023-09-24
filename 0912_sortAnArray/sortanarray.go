package sortanarray

import "vudangkhoa2906-leetcode-go/common"

func sortArray(nums []int) []int {
	len := len(nums)
	common.QuickSort(nums, 0, len-1)
	return nums
}
