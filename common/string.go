package common

func CompareString(str1, str2 string) int {
	len1, len2 := len(str1), len(str2)
	if len1 > len2 {
		return 1
	}
	if len1 < len2 {
		return -1
	}
	for idx := range str1 {
		if str1[idx] > str2[idx] {
			return 1
		}
		if str1[idx] < str2[idx] {
			return -1
		}
	}
	return 0
}

func Reverse(s string) string {
	byteS := []byte(s)
	for startIdx, endIdx := 0, len(s)-1; startIdx < endIdx; startIdx, endIdx = startIdx+1, endIdx-1 {
		byteS[startIdx], byteS[endIdx] = byteS[endIdx], byteS[startIdx]
	}
	return string(byteS)
}

func IsPalindromicSubstring(s string, lowIdx, highIdx int) bool {
	for lowIdx < highIdx {
		if s[lowIdx] != s[highIdx] {
			return false
		}
		lowIdx++
		highIdx--
	}
	return true
}
