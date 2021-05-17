package main

import "fmt"

//go 不支持隐式类型转换
func main() {

	i := 3
	// var f float64 = i  这种隐式类型转换会报错
	var f float64 = float64(i)
	fmt.Println(f) //3

	a := 20
	b := float64(a) / f
	fmt.Println("20/3=", b) //20/3= 6.666666666666667
}
