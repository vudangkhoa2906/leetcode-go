package minimumheighttrees

import (
	"fmt"
	"math"
)

func findMinHeightTrees(n int, edges [][]int) []int {
	resCount := math.MaxInt
	var res []int
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	fmt.Println(adj)
	dp := make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}

	var findHeightBranches func(parentIdx, childIdx int) int
	findHeightBranches = func(parentIdx, childIdx int) int {
		if dp[parentIdx][childIdx] > 0 {
			return dp[parentIdx][childIdx]
		}
		if len(adj[childIdx]) == 1 {
			dp[parentIdx][childIdx] = 1
			return dp[parentIdx][childIdx]
		}
		var resFindHeightBranches int
		for _, otherChildIdx := range adj[childIdx] {
			if otherChildIdx != parentIdx {
				if tempRes := findHeightBranches(childIdx, otherChildIdx); tempRes > resFindHeightBranches {
					resFindHeightBranches = tempRes
				}
			}
		}
		dp[parentIdx][childIdx] = 1 + resFindHeightBranches
		return dp[parentIdx][childIdx]
	}

	var findHeightTrees func(parentIdx int) int
	findHeightTrees = func(parentIdx int) int {
		var resFindHeightTrees int
		for _, childIdx := range adj[parentIdx] {
			if tempRes := findHeightBranches(parentIdx, childIdx); tempRes > resFindHeightTrees {
				resFindHeightTrees = tempRes
			}
		}
		return 1 + resFindHeightTrees
	}

	for idx := 0; idx < n; idx++ {
		if tempRes := findHeightTrees(idx); tempRes < resCount {
			resCount = tempRes
			res = res[:0]
			res = append(res, idx)
		} else if tempRes == resCount {
			res = append(res, idx)
		}
	}

	return res
}
