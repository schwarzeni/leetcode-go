package _0240

// 2020.09.14
//  https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	result := make(chan bool)
	for i := range matrix {
		go func(i int) {
			if matrix[i][0] <= target && matrix[i][len(matrix[0])-1] >= target {
				if binarySearch(matrix[i], target) {
					result <- true
					return
				}
			}
			result <- false
		}(i)
	}
	for i := 0; i < len(matrix); i++ {
		if r := <-result; r {
			return true
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
