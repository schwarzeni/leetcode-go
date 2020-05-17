package _0210

// 2020.05.17
// https://leetcode-cn.com/problems/course-schedule-ii/
// 分别为 BFS 和 DFS 解法
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	degree := make([]int, numCourses)
	var queue []int
	var courseOrder []int
	for _, v := range prerequisites {
		degree[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		courseOrder = append(courseOrder, c)
		for _, nc := range graph[c] {
			degree[nc]--
			if degree[nc] == 0 {
				queue = append(queue, nc)
			}
		}
	}
	if len(courseOrder) == numCourses {
		return courseOrder
	}
	return []int{}
}

//func findOrder(numCourses int, prerequisites [][]int) []int {
//	graph := make(map[int][]int)
//	visited := make([]int, numCourses)
//	var courseOrder []int
//	var dfs func(currCourse int) bool
//	for _, v := range prerequisites {
//		graph[v[1]] = append(graph[v[1]], v[0])
//	}
//	dfs = func(currCourse int) bool {
//		visited[currCourse] = 1
//		for _, c := range graph[currCourse] {
//			if visited[c] == 1 {
//				return false
//			}
//			if visited[c] == 0 {
//				if !dfs(c) {
//					return false
//				}
//			}
//		}
//		visited[currCourse] = 2
//		courseOrder = append(courseOrder, currCourse)
//		return true
//	}
//	for i := 0; i < numCourses; i++ {
//		if visited[i] == 0 {
//			if !dfs(i) {
//				return []int{}
//			}
//		}
//	}
//	i, j := 0, numCourses-1
//	for i < j {
//		courseOrder[i], courseOrder[j] = courseOrder[j], courseOrder[i]
//		i, j = i+1, j-1
//	}
//
//	return courseOrder
//}
