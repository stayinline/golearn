package main

/*
for 循环时go唯一的循环控制语句，也就是没有while循环
*/

func demo1() {
	i := 1
	for i < 3 {
		println(i)
		i++
	}
}

func demo2() {
	for i := 1; i < 3; {
		println(i)
		i++
	}
}

func demo3() {
	for i := 1; i < 3; i++ {
		println(i)
	}
}

func demo4() {
	i := 0
	for {
		if i == 3 {
			break
		}
		println(i)
		i++
	}
}

func demo5() {
	i := 0
	for i != 3 {
		i++
		if i < 3 {
			continue
		}
		println(i)
	}
}

func main() {

	demo1()
	println()
	demo2()
	println()

	demo3()
	println()

	demo4()
	println()

	demo5()
}
