package main

import (
	"fmt"
	"math"
)

//1.全局常量
const s = "constant_string"

//2.用作枚举
const (
	WOMAN   = 0
	MAN     = 1
	UNKNOWN = "qwer"
)

func main() {
	fmt.Println(s)

	//3.局部常量
	const a int = 500000000

	const b = 1234543212345432 / a
	fmt.Println(b)

	fmt.Println(int64(b))

	fmt.Println(math.Sin(float64(a)))

	println(MAN)
	println(UNKNOWN)

}
