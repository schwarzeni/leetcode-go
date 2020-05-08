package _0063

// 2020.05.08
// https://leetcode-cn.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	var rn, cn = len(obstacleGrid), len(obstacleGrid[0])
	count := make([][]int, rn)
	for idx, _ := range count {
		count[idx] = make([]int, cn)
	}
	for idx := 0; idx < cn && obstacleGrid[0][idx] == 0; idx++ {
		count[0][idx] = 1
	}
	for idx := 0; idx < rn && obstacleGrid[idx][0] == 0; idx++ {
		count[idx][0] = 1
	}
	for r := 1; r < rn; r++ {
		for c := 1; c < cn; c++ {
			if obstacleGrid[r][c] == 0 {
				count[r][c] = count[r-1][c] + count[r][c-1]
			}
		}
	}

	return count[rn-1][cn-1]
}
