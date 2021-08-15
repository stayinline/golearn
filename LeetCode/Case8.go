package main

import "math"

func myAtoi(s string) int {
	cur, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		cur = 10*cur + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*cur < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*cur > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * cur
}
func main() {

}
