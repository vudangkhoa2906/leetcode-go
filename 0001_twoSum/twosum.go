package twosum

func twoSum(nums []int, target int) []int {
	var res []int
	indices := make(map[int]int)
	for idx1, val := range nums {
		if idx2, prs := indices[target-val]; !prs {
			indices[val] = idx1
		} else {
			res = make([]int, 2)
			res[0] = idx2
			res[1] = idx1
			return res
		}
	}
	return res
}
