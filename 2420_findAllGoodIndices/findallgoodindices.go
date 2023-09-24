package findallgoodindices

func goodIndices(nums []int, k int) []int {
	lenNums := len(nums)
	nge, pge := make([]int, lenNums), make([]int, lenNums)
	for idx := 1; idx < lenNums; idx++ {
		if nums[idx] <= nums[idx-1] {
			pge[idx] = pge[idx-1] + 1
		} else {

		}
	}
	for idx := lenNums - 2; idx >= 0; idx-- {
		if nums[idx] <= nums[idx+1] {
			nge[idx] = nge[idx+1] + 1
		} else {

		}
	}
	var res []int
	for idx := k; idx < lenNums-k; idx++ {
		if pge[idx-1] >= k-1 && nge[idx+1] >= k-1 {
			res = append(res, idx)
		} else {

		}
	}
	return res
}
