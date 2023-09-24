package deleteoperationfortwostrings

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	resCur, resPrev := make([]int, len2), make([]int, len2)
	var tempCur, tempPrev int
	for _, val1 := range word1 {
		for idx2, val2 := range word2 {
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
			resCur[idx2] = tempCur + tempPrev
		}
		copy(resPrev, resCur)
	}
	return len1 + len2 - 2*resPrev[len2-1]
}
