package onesandzeroes

func findMaxForm(strs []string, m int, n int) int {
	resCur, resPrev := make([][]int, m+1), make([][]int, m+1)
	for idx := 0; idx <= m; idx++ {
		resCur[idx], resPrev[idx] = make([]int, n+1), make([]int, n+1)
	}
	numZeroes, numOnes := countZeroesAndOnes(strs[0])
	for idx1 := 0; idx1 <= m; idx1++ {
		for idx2 := 0; idx2 <= n; idx2++ {
			if idx1 >= numZeroes && idx2 >= numOnes {
				resPrev[idx1][idx2] = 1
			}
		}
	}
	for idx := 1; idx < len(strs); idx++ {
		numZeroes, numOnes = countZeroesAndOnes(strs[idx])
		for idx1 := 0; idx1 <= m; idx1++ {
			for idx2 := 0; idx2 <= n; idx2++ {
				resCur[idx1][idx2] = resPrev[idx1][idx2]
				if idx1 >= numZeroes && idx2 >= numOnes && resCur[idx1][idx2] < 1+resPrev[idx1-numZeroes][idx2-numOnes] {
					resCur[idx1][idx2] = 1 + resPrev[idx1-numZeroes][idx2-numOnes]
				}
			}
		}
		for idx1 := 0; idx1 <= m; idx1++ {
			copy(resPrev[idx1], resCur[idx1])
		}
	}
	return resPrev[m][n]
}

func countZeroesAndOnes(str string) (int, int) {
	var numZeroes, numOnes int
	for idx := 0; idx < len(str); idx++ {
		if str[idx] == '0' {
			numZeroes++
		} else {
			numOnes++
		}
	}
	return numZeroes, numOnes
}
