package palindromepartitioning

import "vudangkhoa2906-leetcode-go/common"

func partition(s string) [][]string {
	var res [][]string
	var bt, tempRes []string
	lenS := len(s)

	var partitionRec func(idx int)
	partitionRec = func(idx int) {
		if idx == lenS {
			tempRes = make([]string, len(bt))
			copy(tempRes, bt)
			res = append(res, tempRes)
			return
		}
		for idx1 := idx; idx1 < lenS; idx1++ {
			if common.IsPalindromicSubstring(s, idx, idx1) {
				bt = append(bt, s[idx:idx1+1])
				partitionRec(idx1 + 1)
				bt = bt[:len(bt)-1]
			}
		}
	}
	partitionRec(0)
	return res
}
