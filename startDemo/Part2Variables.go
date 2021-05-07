package main

import "fmt"

/**
go的变量声明

*/

func main() {

	//隐式声明
	var a = "string_test"
	//显式声明
	var s string = "string_test"
	fmt.Println(a)
	fmt.Println(s)

	//多变量声明
	var b, c int = 11, 22
	fmt.Println(b, c)

	var b2, c2 = 11.2, 22.2
	fmt.Println(b2, c2)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	//省略var，编译器会做自动类型推断，但是 := 左侧必须是一个没有声明过的变量
	f := "hello"
	// f := "hello" 再次声明变量f会报错
	fmt.Println(f)

	//也可以多个使用
	g, h := 123, false
	print(g)
	print(h)

}
