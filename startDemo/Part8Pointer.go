package main

import "fmt"

func main() {

	pointerOne()
	//printSplitLine()

	nullPointerDemo()

	pointerOfArray()

	arrayPointer()

	pointerOfPointer()
}

func pointerOne() {
	a := 10
	fmt.Println("a的值：", a)
	fmt.Println("a的地址：", &a)
	ip := &a
	fmt.Println("ip的值", ip)
	fmt.Println("*ip 的值：", *ip)
}

func nullPointerDemo() {
	var p *int
	fmt.Println("p的值", p)      //p的值 <nil>
	fmt.Printf("p的值: %x\n", p) //p的值: 0

	// 空指针判断
	if p != nil {
		fmt.Println("p 不是空指针")
	}
	if p == nil {
		fmt.Println("p 是空指针")
	}
}

//指针数组---》存放指针的数组
func pointerOfArray() {
	arr := []int{1, 2, 3}
	fmt.Println("数组arr的内容：", arr)

	const LEN int = 3

	//声明一个指针
	var ptrOfArr [LEN]*int
	i := 0
	for ; i < LEN; i++ {
		ptrOfArr[i] = &arr[i]
		fmt.Printf("数组ptrOfArr的内容：ptrOfArr[%d] = 0x%d\n", i, ptrOfArr[i])
	}

	for i = 0; i < LEN; i++ {
		fmt.Printf("数组ptrOfArr的内容：*ptrOfArr[%d] = %d\n", i, *ptrOfArr[i])
	}

}

//数组指针---》存放数组地址的变量 ？？？-> 切片的
func arrayPointer() {
	arr := []int{1, 2, 3}
	ptr := &arr
	fmt.Println("&arr= ", ptr) //&arr=  &[1 2 3]
}

//指针的指针
func pointerOfPointer() {
	i := 100
	var ptr *int     //声明一个指针
	var pPtr **int   //声明一个二级指针
	var ppPtr ***int //声明一个三级指针
	ptr = &i
	pPtr = &ptr
	ppPtr = &pPtr
	fmt.Printf("变量 i 的值= %d, *ptr= %d\n", i, *ptr)         //变量 i 的值= 100, *ptr= 100
	fmt.Printf("变量 i 的值= %d, **pPtr= %d\n", i, **pPtr)     //变量 i 的值= 100, *pPtr= 100
	fmt.Printf("变量 i 的值= %d, ***ppPtr= %d\n", i, ***ppPtr) //变量 i 的值= 100, *ppPtr= 100
}
