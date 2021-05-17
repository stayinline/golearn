package main

import "fmt"

//go把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
///* 定义接口 */
//type interface_name interface {
//   method_name1 [return_type]
//   method_name2 [return_type]
//   method_name3 [return_type]
//   ...
//   method_namen [return_type]
//}
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
//	/* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
//	/* 方法实现*/
//}

type Phone interface {
	call()
	getColor()
}

type IPhone struct {
}

type HuaWeiPhone struct {
}

func (iPhone11 IPhone) call() {
	fmt.Println("iPhone11 calling!!!")
}

func (iPhone11 IPhone) getColor() {
	fmt.Println("iPhone11 is white color")
}

func (mateX HuaWeiPhone) call() {
	fmt.Println("mateX calling!!!")

}

func main() {
	var phone Phone
	//phone.call()
	// 这行代码会报空指针Error：panic: runtime error: invalid memory address or nil pointer dereference

	phone = new(IPhone)
	phone.call()
	phone.getColor()

	mateX := new(HuaWeiPhone)
	mateX.call()

	//iPhone11 calling!!!
	//iPhone11 is white color
	//mateX calling!!!

}
