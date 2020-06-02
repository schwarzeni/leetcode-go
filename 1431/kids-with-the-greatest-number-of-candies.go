package _1431

// 2020.06.01
// https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	result := make([]bool, len(candies))
	for _, v := range candies[1:] {
		if v > max {
			max = v
		}
	}
	for i, v := range candies {
		result[i] = v+extraCandies >= max
	}
	return result
}
