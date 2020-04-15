package _72

// 2020.04.15
// https://leetcode-cn.com/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	} else if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}

	var (
		distance func(word1, word2 string, idx1, idx2 int) int
		cache    = make([][]int, len(word1))
	)
	for idx, _ := range cache {
		cache[idx] = make([]int, len(word2))
		for i, _ := range cache[idx] {
			cache[idx][i] = -1
		}
	}

	distance = func(word1 string, word2 string, idx1, idx2 int) int {
		if word1 == word2 {
			return 0
		} else if len(word1) == 0 {
			return len(word2)
		} else if len(word2) == 0 {
			return len(word1)
		} else if v := cache[idx1][idx2]; v != -1 {
			return v
		}
		var (
			d1  = -1 // 替换
			d2  = -1 // 插入
			d3  = -1 // 删除
			min int
		)
		if d1 = distance(word1[1:], word2[1:], idx1+1, idx2+1); word1[0] != word2[0] {
			d1 += 1
		}

		d2 = distance(word1, word2[1:], idx1, idx2+1) + 1

		d3 = distance(word1[1:], word2, idx1+1, idx2) + 1

		min = d1
		if d3 < d1 {
			min = d3
		}
		if d2 < min {
			min = d2
		}
		if cache[idx1][idx2] == -1 || cache[idx1][idx2] > min {
			cache[idx1][idx2] = min
		}
		return min
	}

	return distance(word1, word2, 0, 0)
}
