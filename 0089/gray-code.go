package _89

// 2020.04.05
// https://leetcode-cn.com/problems/gray-code/

func grayCode(n int) []int {
	var (
		mark         = make(map[string]struct{})
		dataSet      = make([]byte, n)
		result       []int
		changeOneBit = func(idx int) {
			dataSet[idx] = 97 - dataSet[idx]
		}
		toInt = func() (result int) {
			for idx, _ := range dataSet {
				if dataSet[n-1-idx] == '1' {
					result += 1 << idx
				}
			}
			return
		}
	)

	for idx, _ := range dataSet {
		dataSet[idx] = '0'
	}
	mark[string(dataSet)] = struct{}{}
	result = append(result, 0)

LOOP:
	for {
		for idx, _ := range dataSet {
			changeOneBit(idx)
			if _, ok := mark[string(dataSet)]; !ok {
				result = append(result, toInt())
				mark[string(dataSet)] = struct{}{}
				continue LOOP
			}
			changeOneBit(idx)
		}
		break
	}

	return result
}
