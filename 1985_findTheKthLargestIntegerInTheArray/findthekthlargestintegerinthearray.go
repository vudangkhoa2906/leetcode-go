package findthekthlargestintegerinthearray

func kthLargestNumber(nums []string, k int) string {
	heapify(nums)
	var res string
	for i := 0; i < k; i++ {
		res, nums = delMax(nums)
	}
	return res
}

func heapify(nums []string) {
	len := len(nums)
	for parentIdx := len/2 - 1; parentIdx >= 0; parentIdx-- {
		sink(nums, parentIdx, len-1)
	}
}

func exchange(nums []string, idx1, idx2 int) {
	nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
}

func sink(nums []string, parentIdx, highIdx int) {
	var childIdx int
	for 2*parentIdx <= highIdx-1 {
		childIdx = 2*parentIdx + 1
		if childIdx < highIdx && compareString(nums[childIdx], nums[childIdx+1]) < 0 {
			childIdx++
		}
		if compareString(nums[parentIdx], nums[childIdx]) >= 0 {
			break
		}
		exchange(nums, parentIdx, childIdx)
		parentIdx = childIdx
	}
}

func delMax(nums []string) (string, []string) {
	len := len(nums)
	exchange(nums, 0, len-1)
	sink(nums, 0, len-2)
	return nums[len-1], nums[:len-1]
}

func compareString(str1, str2 string) int {
	len1, len2 := len(str1), len(str2)
	if len1 > len2 {
		return 1
	}
	if len1 < len2 {
		return -1
	}
	for idx := range str1 {
		if str1[idx] > str2[idx] {
			return 1
		}
		if str1[idx] < str2[idx] {
			return -1
		}
	}
	return 0
}
