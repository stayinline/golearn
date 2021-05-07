package main

import (
	"fmt"
	"time"
)

/**
这里写狗的switch语句
*/

/*
直接声明一个demo1方法会报错，因为在Part4For中声明了一个: func demo1()
*/
func switchDemo1() {
	i := 2
	println("now,i=", i)

	switch i {
	case 1:
		println("in case1 i=1")
	case 2:
		println("in case2 i=2") //这里展示只匹配一种情况
	case 3:
		println("in case3 1=3")
	}
}

/*
匹配多种情况的时候逗号分隔即可
*/
func switchDemo2() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It`s the weekend")
	default:
		println("It`s a weekday")
	}
}

/*
在case中达到和if语句相同的效果
*/
func switchDemo3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("now is before noon")
	default:
		println("now ia after noon")
	}
}

/*
type-switch 用于判断某个interface的返回值类型
*/
func switchDemo4() {
	var x interface{}

	switch t := x.(type) {
	case int:
		println("x type is int")
	case float64:
		println("x type is float64")

	case func(int):
		println("x type is func(int)")
	default:
		fmt.Println("x type is unknown ,t=", t)
	}
}

func main() {
	switchDemo1()
	switchDemo2()
	switchDemo3()
	switchDemo4()
}
