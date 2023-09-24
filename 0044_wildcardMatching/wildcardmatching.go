package wildcardmatching

func isMatch(s string, p string) bool {
	len1, len2 := len(s), len(p)
	if len2 == 0 {
		return len1 == 0
	}
	if len1 == 0 {
		for _, val2 := range p {
			if val2 != '*' {
				return false
			}
		}
		return true
	}
	resCur, resPrev := make([]bool, len2), make([]bool, len2)
	occurred, matched := false, true
	for idx2 := 0; idx2 < len2; idx2++ {
		if matched {
			if p[idx2] == '*' {
				resPrev[idx2] = true
			} else if p[idx2] == '?' || p[idx2] == s[0] {
				if !occurred {
					occurred = true
					resPrev[idx2] = true
				} else {
					matched = false
				}
			} else {
				matched = false
			}
		}
	}
	var tempPrev, tempCur bool
	for idx1 := 1; idx1 < len1; idx1++ {
		for idx2 := 0; idx2 < len2; idx2++ {
			tempPrev = false
			if s[idx1] == p[idx2] || p[idx2] == '?' {
				tempCur = true
				if idx2 > 0 {
					tempPrev = resPrev[idx2-1]
				}
			} else {
				tempCur = false
				if p[idx2] == '*' {
					tempCur = true
					if idx2 == 0 {
						tempPrev = true
					} else {
						tempPrev = resCur[idx2-1] || resPrev[idx2-1] || resPrev[idx2] // match 0 || match 1 || match > 1
					}
				}
			}
			resCur[idx2] = tempCur && tempPrev
		}
		copy(resPrev, resCur)
	}
	return resPrev[len2-1]
}
