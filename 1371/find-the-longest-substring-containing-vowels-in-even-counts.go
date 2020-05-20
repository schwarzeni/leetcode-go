package _1371

// 2020.05.20
// https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
// 不会 ....
// 对于一个区间，我们可以用两个前缀和的差值，得到其中某个字母的出现次数
// 定义 pre[i][k] 表示在字符串前 i 个字符中，第 k 个元音字母一共出现的次数
// [l, r] --> pre[r][k]−pre[l−1][k]
// 在 O(1) 的时间得到第 k 个元音字母出现的次数
// 对于每一个元音字母，我们都判断一下其是否出现偶数次即可
func findTheLongestSubstring(s string) int {
	pos := make(map[int]int)
	ans, status := 0, 0
	pos[0] = 0
	for i, c := range s {
		switch c {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		v, ok := pos[status]
		if !ok {
			pos[status] = i + 1
		} else if l := i + 1 - v; l > ans {
			ans = l
		}
	}
	return ans
}
