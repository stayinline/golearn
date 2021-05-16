package main

import "fmt"

//声明一个结构体
type Book struct {
	book_name string
	book_id   int
	author    string
}

func main() {
	structDemo1()

	structMemberDemo()

	structAsParam()

}

//结构体当做函数参数
func structAsParam() {
	goBook := Book{
		book_name: "Go程序设计",
		book_id:   123,
		author:    "abcd",
	}
	javaBook := Book{
		book_name: "java程序设计",
		book_id:   124,
		author:    "qwrr",
	}
	cBook := Book{
		book_name: "c程序设计",
		book_id:   125,
		author:    "adsf",
	}
	books := [...]Book{goBook, javaBook, cBook}
	printBookName(books)
	//book[0].book_name= Go程序设计
	//book[1].book_name= java程序设计
	//book[2].book_name= c程序设计

	booksPtr := [...]*Book{&goBook, &javaBook, &cBook}
	printBookNameByPointer(booksPtr)
	//book[0].book_name= Go程序设计
	//book[1].book_name= java程序设计
	//book[2].book_name= c程序设计

}

func structMemberDemo() {
	// 使用member ：value的格式创建一个新的结构体
	goBook := Book{
		book_name: "Go程序设计",
		book_id:   123,
		author:    "abcd",
	}
	fmt.Println("结构体成员变量的访问：goBook.book_name：", goBook.book_name)
	//结构体成员变量的访问：goBook.book_name： Go程序设计
	goBook.book_name = "java程序设计"
	fmt.Println("变更之后：goBook.book_name：", goBook.book_name)
	//变更之后：goBook.book_name： java程序设计
}

func structDemo1() {
	// 使用member ：value的格式创建一个新的结构体
	goBook := Book{
		book_name: "Go程序设计",
		book_id:   123,
		author:    "abcd",
	}
	fmt.Println(goBook)
	//{Go程序设计 123 abcd}
	//只指定某一些字段的值
	goBook2 := Book{
		book_name: "Go程序设计",
	}
	fmt.Println(goBook2)
	//{Go程序设计 0 }
}

//数组做参数需要指定数组长度，这里的3必须写
func printBookName(books [3]Book) {
	i := 0
	for ; i < len(books); i++ {
		fmt.Printf("book[%d].book_name= %s\n", i, books[i].book_name)
	}
}

//结构体指针也是通过.操作符访问成员
func printBookNameByPointer(books [3]*Book) {
	i := 0
	for ; i < len(books); i++ {
		fmt.Printf("book[%d].book_name= %s\n", i, books[i].book_name)
	}
}
