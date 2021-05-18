package main

import (
	"fmt"
	"time"
)

func main() {

	testGoDemo()

	channelDemo()

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

//go channel
// 可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。
// 操作符 <- 用于指定通道的方向，发送或接收。
// 如果未指定方向，则为双向通道。
func channelDemo() {

	//声明一个channel，使用chan关键字
	ch := make(chan int)

	arr := []int{1, 2, 3, 4, 5}
	mid := len(arr) / 2
	go getSumOfArray(ch, arr[:mid])
	go getSumOfArray(ch, arr[mid:])

	x := <-ch
	y := <-ch
	fmt.Printf("x+y=%d\n", x+y)
	//fmt.Println(x)
}

func getSumOfArray(ch chan int, arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	//fmt.Println(sum)
	ch <- sum //将sum发送到通道ch中
}

func getSum2(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

//死锁：
// 只要发送端和接收端有一个阻塞就会死锁
// 参考：https://blog.csdn.net/benben_2015/article/details/84842486
func deadLockDemo() {

}
