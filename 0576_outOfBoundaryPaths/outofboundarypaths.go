package outofboundarypaths

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var res int
	resCur, resPrev := make([][]int, m), make([][]int, m)
	for idx := 0; idx < m; idx++ {
		resCur[idx], resPrev[idx] = make([]int, n), make([]int, n)
	}
	resPrev[startRow][startColumn] = 1
	for move := 1; move <= maxMove; move++ {
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				if r == 0 {
					res = res + resPrev[r][c]
				}
				if r == m-1 {
					res = res + resPrev[r][c]
				}
				if c == 0 {
					res = res + resPrev[r][c]
				}
				if c == n-1 {
					res = res + resPrev[r][c]
				}
				resCur[r][c] = 0
				if r > 0 {
					resCur[r][c] += resPrev[r-1][c]
				}
				if r < m-1 {
					resCur[r][c] += resPrev[r+1][c]
				}
				if c > 0 {
					resCur[r][c] += resPrev[r][c-1]
				}
				if c < n-1 {
					resCur[r][c] += resPrev[r][c+1]
				}
			}
		}
		for r := 0; r < m; r++ {
			copy(resPrev[r], resCur[r])
		}
	}

	return res
}
