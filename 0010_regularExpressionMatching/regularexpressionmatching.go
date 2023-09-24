package regularexpressionmatching

func isMatch(s string, p string) bool {
	len1, len2 := len(s), len(p)
	if len2 == 0 {
		return len1 == 0
	}

	dp := make([][]bool, len1+1)
	for idx1 := 0; idx1 <= len1; idx1++ {
		dp[idx1] = make([]bool, len2)
	}
	var occurred bool
	for idx2 := 1; idx2 < len2; idx2 += 2 {
		if !occurred {
			if p[idx2] == '*' {
				dp[0][idx2] = true
			} else {
				occurred = true
			}
		}
	}

	for idx1 := 1; idx1 <= len1; idx1++ {
		for idx2 := 0; idx2 < len2; idx2++ {
			if s[idx1-1] == p[idx2] || p[idx2] == '.' {
				if idx2 > 0 {
					dp[idx1][idx2] = dp[idx1-1][idx2-1]
				} else {
					dp[idx1][idx2] = idx1 == 1
				}
			} else {
				if p[idx2] == '*' {
					if idx2 < 2 {
						dp[idx1][idx2] = (s[idx1-1] == p[idx2-1] || p[idx2-1] == '.') && dp[idx1-1][idx2] // match >= 1
					} else {
						dp[idx1][idx2] = dp[idx1][idx2-2] || // match 0
							(s[idx1-1] == p[idx2-1] || p[idx2-1] == '.') && dp[idx1-1][idx2] // match >= 1
					}
				}
			}
		}
	}
	return dp[len1][len2-1]
}
