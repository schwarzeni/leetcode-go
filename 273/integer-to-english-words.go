package _273

// 2020.04.11
// https://leetcode-cn.com/problems/integer-to-english-words/
var numToStr = map[int]string{
	0:          "",
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

func parseHundred(num int) string {
	result := ""
	for num > 0 {
		if num > 0 && num <= 20 {
			return result + numToStr[num]
		}
		if num > 20 && num < 100 {
			ten := (num / 10) * 10
			one := num % 10
			if one == 0 {
				return result + numToStr[ten]
			}
			return result + numToStr[ten] + " " + numToStr[one]
		}
		hundred := num / 100
		num = num - hundred*100
		if num%100 == 0 {
			return numToStr[hundred] + " " + numToStr[100]
		}
		result = numToStr[hundred] + " " + numToStr[100] + " "
	}
	return ""
}

func parse(num, divisor int) string {
	if divisor == 1 {
		return parseHundred(num)
	}
	result := ""
	d := num / divisor
	r := num % divisor
	if d > 0 {
		result = parseHundred(d) + " " + numToStr[divisor]
	}
	parseResult := parse(r, divisor/1000)
	if len(parseResult) > 0 && len(result) > 0 {
		result = result + " " + parseResult
	} else {
		result = result + parseResult
	}
	return result
}

func numberToWords(num int) string {
	if num == 0 { // fix case zero failed
		return "Zero"
	}
	return parse(num, 1000000000)
}
