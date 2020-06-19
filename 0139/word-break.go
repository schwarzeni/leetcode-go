package _0139

// 2020.06.19
// https://leetcode-cn.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	predixMap := make(map[int]bool)
	wordDictMap := make(map[string]struct{})

	for i := range wordDict {
		wordDictMap[wordDict[i]] = struct{}{}
	}

	var exploreFunc func(currIdx int) bool
	exploreFunc = func(currIdx int) bool {
		if v, ok := predixMap[currIdx]; ok && !v {
			return false
		}
		predixMap[currIdx] = true
		for i := currIdx; i < len(s); i++ {
			if _, ok := wordDictMap[s[currIdx:i+1]]; ok {
				if i == len(s)-1 {
					return true
				}
				if exploreFunc(i + 1) {
					return true
				}
			}
		}
		predixMap[currIdx] = false
		return false
	}

	return exploreFunc(0)
}
