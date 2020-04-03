package _22

//  2020.04.02
// https://leetcode-cn.com/problems/generate-parentheses/submissions/
func generateParenthesis(n int) []string {
	var (
		result   []string
		generate func(l, r int, str string)
	)

	generate = func(l, r int, str string) {
		if l > n || r > n || r > l {
			return
		}
		if l == n && r == n {
			result = append(result, str)
			return
		}
		generate(l+1, r, str+"(")
		generate(l, r+1, str+")")
	}

	generate(1, 0, "(")

	return result
}
