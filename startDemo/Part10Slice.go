package main

import "fmt"

// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
// Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，
// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
func main() {

	createSliceDemo()

	fourFuncOfSlice()

	sliceCutDemo()

}

//通过s[start:end]的方式截取s中的一部分
func sliceCutDemo() {
	s := make([]int, 3)
	s = append(s, 4, 5, 6)
	s2 := s[2:5]
	fmt.Println("s=", s)
	//s= [0 0 0 4 5 6]
	fmt.Println("2=", s2)
	//2= [0 4 5]
}

//len():求切片的长度
//cap():求切片的容量
//append():向切片中添加元素
//copy():拷贝切片到另一个切片
func fourFuncOfSlice() {
	s := make([]int, 3)
	fmt.Println(len(s))
	//3
	fmt.Println(cap(s))
	//3
	s2 := make([]int, 3, 10)
	fmt.Println(len(s2))
	//3
	fmt.Println(cap(s2))
	//10
	s2 = append(s2, 4, 5, 6)
	fmt.Println("after append: s2= ", s2)
	//after append: s2=  [0 0 0 4 5 6]
	fmt.Println("before copy() s=", s)
	//before copy() s= [0 0 0]
	copy(s, s2)
	fmt.Println("after copy() s=", s)
	//after copy() s= [0 0 0]
	//说明：将s2拷贝到s，显然做了截取，扔掉了长度比s多的部分
	s3 := make([]int, 7)
	copy(s3, s2)
	fmt.Println("s3= ", s3)
	//s3=  [0 0 0 4 5 6 0]
}

func createSliceDemo() {
	//创建切片的语法：slice1 := make([]type, length, capacity),capacity是可选参数
	s1 := make([]int, 3)
	fmt.Println(s1)
	//[0 0 0]
	//通过s1初始化s2
	s2 := s1[0:3]
	fmt.Println(s2)
	//[0 0 0]
}
