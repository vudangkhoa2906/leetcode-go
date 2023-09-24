package common

type SliceElem struct {
	Idx, Val int
}

func IndexOf[V ~int | ~int64](nums []V, element V) int {
	for idx := range nums {
		if nums[idx] == element {
			return idx
		}
	}
	return -1
}

func UpperBound[V ~int | ~int64](nums []V, target V) int {
	len := len(nums)
	lowIdx, highIdx, res := 0, len-1, len
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] > target {
			res, highIdx = midIdx, midIdx-1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return res
}

func LowerBound[V ~int | ~int64](nums []V, target V) int {
	len := len(nums)
	lowIdx, highIdx, res := 0, len-1, -1
	var midIdx int
	for lowIdx <= highIdx {
		midIdx = (lowIdx + highIdx) / 2
		if nums[midIdx] >= target {
			res, highIdx = midIdx, midIdx-1
		} else {
			lowIdx = midIdx + 1
		}
	}
	return res
}

func NextGreaterElement(nums []int) []int {
	lenNums := len(nums)
	res := make([]int, lenNums)
	stack := make([]SliceElem, 0, lenNums)
	for idx := lenNums - 1; idx >= 0; idx-- {
		for count := len(stack); count > 0 && stack[count-1].Val <= nums[idx]; count = len(stack) {
			stack = stack[:count-1]
		}
		if count := len(stack); count > 0 {
			res[idx] = stack[count-1].Idx
		} else {
			res[idx] = lenNums
		}
		stack = append(stack, SliceElem{Idx: idx, Val: nums[idx]})
	}
	return res
}

func PreviousGreaterElement(nums []int) []int {
	lenNums := len(nums)
	res := make([]int, lenNums)
	stack := make([]SliceElem, 0, lenNums)
	for idx := 0; idx < lenNums; idx++ {
		for count := len(stack); count > 0 && stack[count-1].Val <= nums[idx]; count = len(stack) {
			stack = stack[:count-1]
		}
		if count := len(stack); count > 0 {
			res[idx] = stack[count-1].Idx
		} else {
			res[idx] = -1
		}
		stack = append(stack, SliceElem{Idx: idx, Val: nums[idx]})
	}
	return res
}
