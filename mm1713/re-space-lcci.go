package mm1713

import "math"

// 2020.07.09
// https://leetcode-cn.com/problems/re-space-lcci/
func respace(dictionary []string, sentence string) int {
	dictMap := make(map[string]struct{})
	for i := range dictionary {
		dictMap[dictionary[i]] = struct{}{}
	}
	minUndiscern := make(map[int]int)
	slen := len(sentence)

	var scanSentence func(startIdx int) int
	scanSentence = func(startIdx int) int {
		if startIdx >= slen {
			return 0
		}
		if v, ok := minUndiscern[startIdx]; ok {
			return v
		}
		min := math.MaxInt32
		for i := startIdx; i < slen; i++ {
			ss := sentence[startIdx : i+1]
			l := 0
			if _, ok := dictMap[ss]; !ok {
				l = i - startIdx + 1
			}
			if l += scanSentence(i + 1); l < min {
				min = l
			}
		}
		minUndiscern[startIdx] = min
		return min
	}
	return scanSentence(0)
}
