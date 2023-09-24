package wordbreak

func wordBreak(s string, wordDict []string) bool {
	lenS := len(s)
	dp := make([]bool, lenS+1)
	dp[0] = true
	var lenWord int

	for idx := 1; idx <= lenS; idx++ {
		for _, word := range wordDict {
			lenWord = len(word)
			if idx >= lenWord {
				if s[idx-lenWord:idx] == word && dp[idx-lenWord] {
					dp[idx] = true
					break
				}
			}
		}
	}

	return dp[lenS]
}
