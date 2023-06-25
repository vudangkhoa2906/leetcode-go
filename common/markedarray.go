package common

func MarkedArray(numRows, numCols int) [][]bool {
	res := make([][]bool, numRows)
	for r := 0; r < numRows; r++ {
		res[r] = make([]bool, numCols)
	}
	return res
}
