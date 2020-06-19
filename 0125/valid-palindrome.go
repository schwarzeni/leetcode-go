package _0125

// 2020.06.19
// https://leetcode-cn.com/problems/valid-palindrome/
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for {
		for i < j && !isValid(s[i]) {
			i++
		}
		for i < j && !isValid(s[j]) {
			j--
		}
		if i >= j {
			return true
		}
		if !isEqual(s[i], s[j]) {
			return false
		}
		i++
		j--
	}
}

func isValid(s uint8) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9')
}

func isEqual(s1 uint8, s2 uint8) bool {
	if s1 >= 'a' && s1 <= 'z' {
		s1 = s1 - ('z' - 'Z')
	}
	if s2 >= 'a' && s2 <= 'z' {
		s2 = s2 - ('z' - 'Z')
	}
	return s1 == s2
}
