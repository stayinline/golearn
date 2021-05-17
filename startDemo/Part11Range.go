package main

import "fmt"

//range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
// 在数组和切片中它返回元素的索引和索引对应的值
// 在集合中返回 key-value 对。
func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0
	//（1）当我们不需要知道每个元素的其索引的时候：_
	for _, num := range nums {
		sum += num
	}
	fmt.Println("the sum of all nums is ", sum) //6

	//（2）当我们需要知道每个元素及其索引的时候：
	for i, num := range nums {
		fmt.Printf("nums[%d] = %d\n", i, num)
	}
	map1 := map[int]string{1: "aa", 2: "bb"}
	fmt.Println(map1) //map[1:aa 2:bb]
	for k, v := range map1 {
		fmt.Printf("%d -> %s\n", k, v)
	}
	//1 -> aa
	//2 -> bb

	//（3）range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	//0 103
	//1 111
}
