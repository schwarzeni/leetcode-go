package _0126

// 2020.06.07
// https://leetcode-cn.com/problems/word-ladder-ii/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	graph := make([]map[int]struct{}, len(wordList))
	flag := false
	for i, _ := range wordList {
		if wordList[i] == endWord {
			flag = true
		}
		graph[i] = make(map[int]struct{})
	}
	if flag == false {
		return nil
	}
	for i, _ := range wordList {
		for ii, _ := range wordList {
			_, ok := graph[ii][i]
			if ok || isNear(wordList[i], wordList[ii], 1) {
				graph[i][ii] = struct{}{}
			}
		}
	}
	var path []string
	var result [][]string
	var visited = make([]bool, len(wordList))
	var dfs func(i int)
	minLen := -1
	dfs = func(i int) {
		//visited[i] = true
		defer func() {
			//visited[i] = false
			path = path[:len(path)-1]
		}()
		path = append(path, wordList[i])
		if wordList[i] == endWord {
			r := make([]string, len(path))
			copy(r, path)
			if minLen == -1 || minLen > len(r) {
				minLen = len(r)
				result = [][]string{r}
			} else if minLen == len(r) {
				result = append(result, r)
			}
			return
		}
		for nextI, _ := range graph[i] {
			if !visited[nextI] && (minLen == -1 || len(path) <= minLen-1) {
				visited[nextI] = true
				dfs(nextI)
				visited[nextI] = false
			}
		}
	}
	path = append(path, beginWord)
	for i, _ := range wordList {
		if isNear(beginWord, wordList[i], 1) {
			visited[i] = true
			dfs(i)
			visited[i] = false
		}
	}
	return result
}
func isNear(str1, str2 string, l int) bool {
	for i, _ := range str1 {
		if str1[i] != str2[i] {
			return str1[i+1:] == str2[i+1:]
		}
	}
	return false
}

//func isNear(str1, str2 string, l int) bool {
//	var sameCount int
//	for i, _ := range str1 {
//		if str1[i] != str2[i] {
//			sameCount++
//			if sameCount > l {
//				return false
//			}
//		}
//	}
//	return sameCount == l
//}
