package main

import "fmt"

func main() {

	a, b := getTop2(1, 2, 3)
	fmt.Println(a, b)
}

/**
函数样例

注意：这里的 param1, parm2是值传递
*/
func funcName(param1 int, parm2 int) int {
	return 0
}

//这里的 param1, parm2是引用传递
func funcRef(param1 *int, parm2 *int) int {
	return 0
}

//返回较大值
func getMax(a1 int, a2 int) int {
	if a1 > a2 {
		return a1
	} else {
		return a2
	}
}

//go支持返回多个值
func getTop2(a1 int, a2 int, a3 int) (int, int) {
	first := 0
	second := 0
	return first, second
}
