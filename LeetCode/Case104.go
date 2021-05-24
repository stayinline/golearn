package main

import "fmt"

//func maxDepth(root *TreeNode) int {
//	if nil == root {
//		return 0
//	}
//	depL := maxDepth(root.Left)
//	depR := maxDepth(root.Right)
//	return getMax(depL, depR) + 1
//}
func main() {
	//maxDepth()

	test1()
	test2()
}

func test1() {
	m := []int{1, 2, 3}
	m2 := m
	//m2 = append(m2, 4)
	m2[0] = 100
	fmt.Println(m[0]) // ?
}

func test2() {
	m := []int{1, 2, 3}
	m2 := m
	m2 = append(m2, 4)
	m2[0] = 100
	fmt.Println(m[0]) // ?
}
