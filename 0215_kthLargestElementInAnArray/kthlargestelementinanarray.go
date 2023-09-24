package kthlargestelementinanarray

import "vudangkhoa2906-leetcode-go/common"

func findKthLargest(nums []int, k int) int {
	common.Heapify(nums)
	var res int
	for i := 0; i < k; i++ {
		res, nums = common.DelMax(nums)
	}
	return res
}
