package _0680

// 2020.05.19
// https://leetcode-cn.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	i, j, ok := check(s)
	if ok {
		return true
	}
	_, _, ok1 := check(s[i+1 : j+1])
	_, _, ok2 := check(s[i:j])
	return ok1 || ok2
}

func check(s string) (i, j int, ok bool) {
	i, j = 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return i, j, false
		}
		i, j = i+1, j-1
	}
	return i, j, true
}
