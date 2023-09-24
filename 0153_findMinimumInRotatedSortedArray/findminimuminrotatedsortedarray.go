package findminimuminrotatedsortedarray

import "math"

func findMin(nums []int) int {
	len := len(nums)
	lowIdx, highIdx, res := 0, len-1, math.MaxInt
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[lowIdx] <= nums[midIdx] {
			if res > nums[lowIdx] {
				res = nums[lowIdx]
			}
			lowIdx = midIdx + 1
		} else {
			if res > nums[midIdx] {
				res = nums[midIdx]
			}
			highIdx = midIdx - 1
		}
	}
	return res
}
