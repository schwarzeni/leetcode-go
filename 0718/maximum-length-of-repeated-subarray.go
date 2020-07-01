package _0718

// 2020.07.01
// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
func findLength(A []int, B []int) int {
	dp, maxL := make([][]int, len(A)), 0
	for i := range dp {
		dp[i] = make([]int, len(B))
	}
	for i := range dp {
		if A[i] == B[0] {
			dp[i][0] = 1
			maxL = 1
		}
	}
	for i := range dp[0] {
		if B[i] == A[0] {
			dp[0][i] = 1
			maxL = 1
		}
	}
	for idxA := 1; idxA < len(A); idxA++ {
		for idxB := 1; idxB < len(B); idxB++ {
			if A[idxA] == B[idxB] {
				dp[idxA][idxB] = dp[idxA-1][idxB-1] + 1
				if dp[idxA][idxB] > maxL {
					maxL = dp[idxA][idxB]
				}
			}
		}
	}
	return maxL
}
