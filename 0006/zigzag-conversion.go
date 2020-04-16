package _6

// 2020.04.09
// https://leetcode-cn.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	var (
		sLen   = len(s)
		result = make([]byte, sLen)
		gap    = numRows*2 - 2
		idx    = 0
	)
	if numRows <= 1 {
		return s
	}
	for r := 0; r < numRows; r++ {
		g := gap - r*2
		for c := r; c < sLen; c += gap {
			if idx < sLen {
				result[idx] = s[c]
				idx += 1
			}
			i := c + g
			if i < sLen && i < c+gap && i > c {
				result[idx] = s[i]
				idx += 1
			}
		}
	}

	return string(result)
}

// L    I
// E   ES  G
// E  D H N
// T O  II
// C    R
// n n-2
// one part = 2n-2
// 0     8
// 1    79      15
// 2   6 10   14
// 3  5  11 13
// 4     12
