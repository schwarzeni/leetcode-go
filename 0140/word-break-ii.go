package _0140

import "strings"

// 2020.06.19
// https://leetcode-cn.com/problems/word-break-ii/
func wordBreak(s string, wordDict []string) []string {
	predixMap := make(map[int]bool)
	wordDictMap := make(map[string]struct{})

	for i := range wordDict {
		wordDictMap[wordDict[i]] = struct{}{}
	}

	var res, curr []string

	var exploreFunc func(currIdx int) bool
	exploreFunc = func(currIdx int) bool {
		if v, ok := predixMap[currIdx]; ok && !v {
			return false
		}
		predixMap[currIdx] = true
		flag := false
		for i := currIdx; i < len(s); i++ {
			if _, ok := wordDictMap[s[currIdx:i+1]]; ok {
				curr = append(curr, s[currIdx:i+1])
				if i == len(s)-1 {
					res = append(res, strings.Join(curr, " "))
					flag = true
				} else if exploreFunc(i + 1) {
					flag = true
				}
				curr = curr[:len(curr)-1]
			}
		}
		predixMap[currIdx] = flag
		return flag
	}

	exploreFunc(0)
	return res
}
