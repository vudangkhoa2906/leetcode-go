package findgooddaystorobthebank

func goodDaysToRobBank(security []int, time int) []int {
	lenSecurity := len(security)
	nge, pge := make([]int, lenSecurity), make([]int, lenSecurity)
	for idx := 1; idx < lenSecurity; idx++ {
		if security[idx] <= security[idx-1] {
			pge[idx] = pge[idx-1] + 1
		}
	}
	for idx := lenSecurity - 2; idx >= 0; idx-- {
		if security[idx] <= security[idx+1] {
			nge[idx] = nge[idx+1] + 1
		}
	}
	var res []int
	for idx := time; idx < lenSecurity-time; idx++ {
		if nge[idx] >= time && pge[idx] >= time {
			res = append(res, idx)
		}
	}
	return res
}
