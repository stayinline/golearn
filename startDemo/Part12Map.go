package main

import "fmt"

func main() {

	var mytestMap map[int]string
	fmt.Println("只声明不赋值的map,初始值是：", mytestMap)
	if nil == mytestMap {
		fmt.Println("mytestMap 是 nil")
	}

	mytestMap = make(map[int]string)
	fmt.Println("只声明不赋值的map,初始值是：", mytestMap)

	if nil != mytestMap {
		fmt.Println("mytestMap 不是 nil")
	}

	mytestMap[0] = "java"
	mytestMap[1] = "go"
	mytestMap[2] = "c/c++"

	for k, v := range mytestMap {
		fmt.Printf("%d -> %s\n", k, v)
	}
	//0 -> java
	//1 -> go
	//2 -> c/c++

	// map[k] 返回值有两个，第一个是k对应的val，第二个是k是否存在
	val, ret := mytestMap[2]
	fmt.Println(val) //c/c++
	fmt.Println(ret) //true

	val2, ret2 := mytestMap[6]
	if ret2 {
		fmt.Println("k 为 6 时对应的val=", val2) //不会输出这个
	} else {
		fmt.Println("k 为 6 时对应的 val 不存在") //会输出这个
	}

	//delete函数
	delete(mytestMap, 2)
	fmt.Println(mytestMap) //map[0:java 1:go]

}
