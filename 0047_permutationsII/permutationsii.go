package permutationsii

func permuteUnique(nums []int) [][]int {
	freq := make(map[int]int)
	for _, val := range nums {
		if _, prs := freq[val]; !prs {
			freq[val] = 1
		} else {
			freq[val]++
		}
	}
	keyFreq := make([]int, 0, len(freq))
	for key := range freq {
		keyFreq = append(keyFreq, key)
	}
	var res [][]int
	var bt []int

	var permuteUniqueRec func()
	permuteUniqueRec = func() {
		if lenBt := len(bt); lenBt == len(nums) {
			tempRes := make([]int, lenBt)
			copy(tempRes, bt)
			res = append(res, tempRes)
		} else {
			for _, val := range keyFreq {
				if freq[val] > 0 {
					freq[val]--
					bt = append(bt, val)
					permuteUniqueRec()
					bt = bt[:len(bt)-1]
					freq[val]++
				}
			}
		}
	}

	permuteUniqueRec()
	return res
}
