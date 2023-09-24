package shortestcommonsupersequence

import "strings"

func shortestCommonSupersequence(str1 string, str2 string) string {
	len1, len2 := len(str1), len(str2)
	resCur, resPrev := make([]*strings.Builder, len2), make([]*strings.Builder, len2)
	var tempPrev string
	for idx1, val1 := range str1 {
		for idx2, val2 := range str2 {
			resCur[idx2] = &strings.Builder{}
			if val1 == val2 {
				tempPrev = ""
				if idx2 > 0 && idx1 > 0 {
					tempPrev = resPrev[idx2-1].String()
				}
				resCur[idx2].WriteString(tempPrev)
				resCur[idx2].WriteRune(val2)
			} else {
				tempPrev = ""
				if idx1 > 0 {
					tempPrev = resPrev[idx2].String()
				}
				if idx2 > 0 && resCur[idx2-1].Len() > len(tempPrev) {
					tempPrev = resCur[idx2-1].String()
				}
				resCur[idx2].WriteString(tempPrev)
			}
		}
		copy(resPrev, resCur)
	}
	startIdx1, startIdx2 := 0, 0
	sB := strings.Builder{}
	for _, val := range resPrev[len2-1].String() {
		for idx1 := startIdx1; idx1 < len1; idx1++ {
			if str1[idx1] == uint8(val) {
				startIdx1 = idx1 + 1
				break
			}
			sB.WriteByte(str1[idx1])
		}
		for idx2 := startIdx2; idx2 < len2; idx2++ {
			if str2[idx2] == uint8(val) {
				startIdx2 = idx2 + 1
				break
			}
			sB.WriteByte(str2[idx2])
		}
		sB.WriteByte(byte(val))
	}
	sB.WriteString(str1[startIdx1:])
	sB.WriteString(str2[startIdx2:])
	return sB.String()
}
