package singlenumber

func singleNumber(nums []int) int {
	var res int
	for _, val := range nums {
		res ^= val
	}
	return res
}
