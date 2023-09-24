package common

func HeapSort(nums []int) {
	Heapify(nums)
	highIdx := len(nums) - 1
	for highIdx > 0 {
		Exchange(nums, 0, highIdx)
		highIdx--
		Sink(nums, 0, highIdx)
	}
}

func Heapify(nums []int) {
	len := len(nums)
	for parentIdx := len/2 - 1; parentIdx >= 0; parentIdx-- {
		Sink(nums, parentIdx, len-1)
	}
}

func Sink(nums []int, parentIdx, highIdx int) {
	var childIdx int
	for 2*parentIdx+1 <= highIdx {
		childIdx = 2*parentIdx + 1
		if childIdx < highIdx && nums[childIdx] < nums[childIdx+1] {
			childIdx++
		}
		if nums[parentIdx] > nums[childIdx] {
			break
		}
		Exchange(nums, parentIdx, childIdx)
		parentIdx = childIdx
	}
}

func Exchange(nums []int, idx1, idx2 int) {
	nums[idx1], nums[idx2] = nums[idx2], nums[idx1]
}

func DelMax(nums []int) (int, []int) {
	len := len(nums)
	Exchange(nums, 0, len-1)
	Sink(nums, 0, len-2)
	return nums[len-1], nums[:len-1]
}
