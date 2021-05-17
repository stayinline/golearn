package main

import (
	"fmt"
	"time"
)

func main() {

	testGoDemo()

}

//创建轻量级线程的语法格式：go 函数名( 参数列表 )
func testGoDemo() {
	go say("hello")
	say("world")
	//"hello"和"world"交替打印，并且不存在一个固定的先后顺序
}

func say(s string) {
	for i := 0; i < 4; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
