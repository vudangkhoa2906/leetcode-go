package distinctsubsequences

func numDistinct(s string, t string) int {
	len1, len2 := len(s), len(t)
	resCur, resPrev := make([]int, len2), make([]int, len2)
	if s[0] == t[0] {
		resPrev[0] = 1
	}
	for idx1 := 1; idx1 < len1; idx1++ {
		for idx2 := 0; idx2 < len2; idx2++ {
			resCur[idx2] = resPrev[idx2]
			if s[idx1] == t[idx2] {
				if idx2 > 0 {
					resCur[idx2] += resPrev[idx2-1]
				} else {
					resCur[idx2] += 1
				}
			}
		}
		copy(resPrev, resCur)
	}
	return resPrev[len2-1]
}
