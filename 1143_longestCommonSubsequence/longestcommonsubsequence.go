package longestcommonsubsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	len2 := len(text2)
	resCur, resPrev := make([]int, len2), make([]int, len2)
	var tempCur, tempPrev int
	for _, val1 := range text1 {
		for idx2, val2 := range text2 {
			if val1 == val2 {
				tempCur = 1
				tempPrev = 0
				if idx2 > 0 {
					tempPrev = resPrev[idx2-1]
				}
			} else {
				tempCur = 0
				tempPrev = resPrev[idx2]
				if idx2 > 0 && resCur[idx2-1] > tempPrev {
					tempPrev = resCur[idx2-1]
				}
			}
			resCur[idx2] = tempPrev + tempCur
		}
		copy(resPrev, resCur)
	}
	return resPrev[len2-1]
}
