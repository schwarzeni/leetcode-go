package _0074

// 2020.05.19
// https://leetcode-cn.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	rc := len(matrix)
	if rc == 0 || len(matrix[0]) == 0 || matrix[0][0] > target || matrix[rc-1][len(matrix[0])-1] < target {
		return false
	}

	r := 0
	for ; r < rc && matrix[r][0] <= target; r++ {
	}
	r--
	if matrix[r][0] == target {
		return true
	}

	i, j := 0, len(matrix[0])-1
	for i <= j {
		mid := i + (j-i)/2
		if matrix[r][mid] == target {
			return true
		} else if matrix[r][mid] > target {
			j = mid - 1
		} else if matrix[r][mid] < target {
			i = mid + 1
		}
	}
	return false
}
