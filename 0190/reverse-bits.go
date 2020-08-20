package _0190

// 2020.08.20
// https://leetcode-cn.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 32; i > 0; i-- {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}
	return ans
}
