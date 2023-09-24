package wordsearch

func exist(board [][]byte, word string) bool {
	numRows, numCols := len(board), len(board[0])
	len := len(word)

	var depthFirstSearch func(int, int, int) bool

	depthFirstSearch = func(r, c, idx int) bool {
		if idx == len {
			return true
		}
		if r < 0 || r >= numRows || c < 0 || c >= numCols || board[r][c] != word[idx] {
			return false
		}
		char := board[r][c]
		board[r][c] = 0
		defer func() {
			board[r][c] = char
		}()
		return depthFirstSearch(r+1, c, idx+1) || depthFirstSearch(r-1, c, idx+1) ||
			depthFirstSearch(r, c-1, idx+1) || depthFirstSearch(r, c+1, idx+1)
	}

	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if depthFirstSearch(r, c, 0) {
				return true
			}
		}
	}
	return false
}
