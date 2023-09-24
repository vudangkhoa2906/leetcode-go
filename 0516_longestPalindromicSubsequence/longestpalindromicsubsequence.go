package longestpalindromicsubsequence

func longestPalindromeSubseq(s string) int {
	len := len(s)
	resCur, resPrev := make([]int, len), make([]int, len)
	var tempCur, tempPrev int
	for idx1 := range s {
		for idx2 := len - 1; idx2 >= 0; idx2-- {
			if s[idx1] == s[idx2] {
				tempCur = 1
				tempPrev = 0
				if idx2 < len-1 {
					tempPrev = resPrev[idx2+1]
				}
			} else {
				tempCur = 0
				tempPrev = resPrev[idx2]
				if idx2 < len-1 && resCur[idx2+1] > tempPrev {
					tempPrev = resCur[idx2+1]
				}
			}
			resCur[idx2] = tempCur + tempPrev
		}
		copy(resPrev, resCur)
	}
	return resPrev[0]
}
