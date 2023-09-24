package uncrossedlines

import "vudangkhoa2906-leetcode-go/common"

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	size1, size2 := len(nums1), len(nums2)
	resCur, resPrev := make([]int, size2+1), make([]int, size2+1)
	for idx1 := 0; idx1 < size1; idx1++ {
		for idx2 := 0; idx2 < size2; idx2++ {
			if nums1[idx1] == nums2[idx2] {
				resCur[idx2+1] = 1 + resPrev[idx2]
			} else {
				resCur[idx2+1] = common.Max(resCur[idx2], resPrev[idx2+1])
			}
		}
		copy(resPrev, resCur)
	}
	return resPrev[size2]
}
