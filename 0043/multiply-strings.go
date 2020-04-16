package _43

import "fmt"

// 2020.04.13
// https://leetcode-cn.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var grid = make([][]byte, len(num2))
	for idx, _ := range grid {
		grid[idx] = make([]byte, len(num1)+1)
	}

	add := 0
	for r := 0; r < len(num2); r++ {
		add = 0
		for c := len(num1) - 1; c >= 0; c-- {
			result := int(num1[c]-48)*int(num2[len(num2)-r-1]-48) + add
			grid[r][c+1] = byte(result%10 + 48)
			add = result / 10
		}
		grid[r][0] = byte(add + 48)
	}

	add = 0
	var resultArr []byte
	for c := len(num1); c >= 0; c-- {
		r := 0
		cc := c
		result := add
		for {
			result += int(grid[r][cc] - 48)
			r += 1
			cc += 1
			if r >= len(num2) || cc > len(num1) {
				break
			}
		}
		resultArr = append(resultArr, byte(result%10+48))
		add = result / 10
	}

	for r := 1; r < len(num2); r++ {
		rr := r
		cc := 0
		result := add
		for {
			result += int(grid[rr][cc] - 48)
			rr += 1
			cc += 1
			if rr >= len(num2) || cc > len(num1) {
				break
			}
		}
		resultArr = append(resultArr, byte(result%10+48))
		add = result / 10
	}

	i := 0
	j := len(resultArr) - 1
	for j >= 0 {
		if resultArr[j] == '0' {
			j -= 1
		} else {
			resultArr = resultArr[0 : j+1]
			break
		}
	}
	for i < j {
		resultArr[i], resultArr[j] = resultArr[j], resultArr[i]
		i += 1
		j -= 1
	}
	return string(resultArr)
}

//              1 2 3
//              4 5 6
//              -----
//              7 3 8
//            6 1 5
//        0 4 9 2
//
//            0 7 3 8
//            0 6 1 5
//            0 4 9 2

//          1 2 3 4 5 6 7
//                    1 1
//        0 1 2 3 4 5 6 7
//      0 1 2 3 4 5 6 7

//        9 9 9
//        8 8 8
//      7 9 9 2
//    7 9 9 2
//  7 9 9 2

func printGrid(data [][]byte) {
	for _, cs := range data {
		for _, v := range cs {
			fmt.Print(string(v), " ")
		}
		fmt.Println()
	}
}
