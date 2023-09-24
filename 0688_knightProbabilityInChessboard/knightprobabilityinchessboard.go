package knightprobabilityinchessboard

func knightProbability(n int, k int, row int, column int) float64 {
	var res float64
	resCur, resPrev := make([][]float64, n), make([][]float64, n)
	for idx := 0; idx < n; idx++ {
		resCur[idx], resPrev[idx] = make([]float64, n), make([]float64, n)
	}
	resPrev[row][column] = 1
	for move := 1; move <= k; move++ {
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				resCur[r][c] = 0
				if r >= 2 && c <= n-2 {
					resCur[r][c] += resPrev[r-2][c+1] * 0.125
				}
				if r >= 1 && c <= n-3 {
					resCur[r][c] += resPrev[r-1][c+2] * 0.125
				}
				if r <= n-2 && c <= n-3 {
					resCur[r][c] += resPrev[r+1][c+2] * 0.125
				}
				if r <= n-3 && c <= n-2 {
					resCur[r][c] += resPrev[r+2][c+1] * 0.125
				}
				if r <= n-3 && c >= 1 {
					resCur[r][c] += resPrev[r+2][c-1] * 0.125
				}
				if r <= n-2 && c >= 2 {
					resCur[r][c] += resPrev[r+1][c-2] * 0.125
				}
				if r >= 1 && c >= 2 {
					resCur[r][c] += resPrev[r-1][c-2] * 0.125
				}
				if r >= 2 && c >= 1 {
					resCur[r][c] += resPrev[r-2][c-1] * 0.125
				}
			}
		}
		for r := 0; r < n; r++ {
			copy(resPrev[r], resCur[r])
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			res += resPrev[r][c]
		}
	}

	return res
}
