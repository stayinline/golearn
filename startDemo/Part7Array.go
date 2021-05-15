package main

import "fmt"

func main() {

	//声明并初始化一个数组
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1) //[1 2 3]

	//用下标初始化数组
	arr2 := [5]int{1: 10, 2, 3, 4: 100}
	fmt.Println(arr2) //[0 10 2 3 100]

	//不确定数组长度时候使用...
	arr3 := [...]int{1, 2}
	fmt.Println(arr3) //[1 2]

	//通过下标访问数组
	a1 := arr1[1]
	fmt.Println(a1) //2

	arr21 := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Println(arr21) //[[0 1 2 3] [4 5 6 7] [8 9 10 11]]

	//声明一个二维数组
	var arr22 [][]int
	//arr22 = append(arr22, arr1)
	//arr22 = append(arr22, arr1)
	fmt.Println(arr22) //[1 2 3]

}
