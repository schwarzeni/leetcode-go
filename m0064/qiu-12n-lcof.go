package m0064

// 2020.06.02
//https://leetcode-cn.com/problems/qiu-12n-lcof/
func sumNums(n int) int {
	//ans := 0
	//var sum func(int) bool
	//sum = func(n int) bool {
	//	ans += n
	//	return n > 0 && sum(n-1)
	//}
	//sum(n)
	//return ans
	return quickMulti((1+n), n) >> 1

}

// 「俄罗斯农民乘法」
func quickMulti(A, B int) int {
	ans := 0
	for ; B > 0; B >>= 1 {
		if B&1 > 0 {
			ans += A
		}
		A <<= 1
	}
	return ans
}
