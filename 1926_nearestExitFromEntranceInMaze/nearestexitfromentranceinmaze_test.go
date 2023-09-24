package nearestexitfromentranceinmaze

import "testing"

func TestNearestExitFromEntranceInMaze(t *testing.T) {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	entrance := []int{1, 2}
	if res := nearestExit(maze, entrance); res != 1 {
		t.Errorf("Output: %d: Expected: %d\n", res, 1)
	}
}
