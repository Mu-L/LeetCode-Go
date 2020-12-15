package leetcode

import (
	"strconv"
)

func addFive(n int) int {
	if n > 0 {
		str, index := strconv.Itoa(n), 0
		for index = 0; index < len(str); index++ {
			if int(str[index]-'0') < 5 {
				break
			}
		}
		res, _ := strconv.Atoi(str[:index] + "5" + str[index:])
		return res
	}
	str, index := strconv.Itoa(-n), 0
	for index = 0; index < len(str); index++ {
		if int(str[index]-'0') > 5 {
			break
		}
	}
	res, _ := strconv.Atoi(str[:index] + "5" + str[index:])
	return -res
}

// func addFirstPosition(n int, isPositive bool) (int, bool) {
// 	str, res, flag := strconv.Itoa(n), 0, true
// 	first := n / int(math.Pow10(len(str)-1))
// 	if isPositive {
// 		if first > 5 {
// 			res = first*int(math.Pow10(len(str))) + 5*int(math.Pow10(len(str)-1)) + n%int(math.Pow10(len(str)-1))
// 			flag = true
// 		} else {
// 			flag = false
// 		}
// 	} else {
// 		if first > 5 {
// 			res = -5*int(math.Pow10(len(str))) + n
// 			flag = true
// 		} else {
// 			flag = false
// 		}
// 	}
// 	return res, flag
// }
