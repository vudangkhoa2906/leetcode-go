package searchinputposition

func searchInsert(nums []int, target int) int {
	len := len(nums)
	lowIdx, highIdx, res := 0, len-1, len
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] >= target {
			res = midIdx
			highIdx = midIdx - 1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return res
}
