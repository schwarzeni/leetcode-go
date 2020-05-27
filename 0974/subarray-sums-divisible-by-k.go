package _0974

// 2020.05.27
// https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/
func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}

//func subarraysDivByK(A []int, K int) int {
//	count := 0
//	cache := make(map[int]int)
//	for _, v := range A {
//		c := make(map[int]int)
//		c[v] = 1
//		for k, _ := range cache {
//			c[k+v] += cache[k]
//		}
//		for k, v := range c {
//			if k%K == 0 {
//				count += v
//			}
//		}
//		cache = c
//	}
//	return count
//}

//func subarraysDivByK(A []int, K int) int {
//	count := 0
//	for i := 0; i < len(A); i++ {
//		sum := 0
//		for j := i; j < len(A); j++ {
//			if sum += A[j]; sum%K == 0 {
//				count++
//			}
//		}
//	}
//	return count
//}
