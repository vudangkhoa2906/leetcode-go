package maximumlengthofrepeatedsubarray

func findLength(nums1 []int, nums2 []int) int {
	lenNums1, lenNums2 := len(nums1), len(nums2)
	dp := make([][]int, lenNums1)
	var res int
	for idx := range dp {
		dp[idx] = make([]int, lenNums2)
	}
	for idx1 := 0; idx1 < lenNums1; idx1++ {
		for idx2 := 0; idx2 < lenNums2; idx2++ {
			if nums1[idx1] == nums2[idx2] {
				if idx1 > 0 && idx2 > 0 {
					dp[idx1][idx2] = 1 + dp[idx1-1][idx2-1]
				} else {
					dp[idx1][idx2] = 1
				}
			}
			if res < dp[idx1][idx2] {
				res = dp[idx1][idx2]
			}
		}
	}
	return res
}
