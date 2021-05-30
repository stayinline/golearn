package main

import (
	"fmt"
	"math"
)

//方法1：全局变量法

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	outOrderHelper(root)
	return res
}

func outOrderHelper(root *TreeNode) int {
	if nil == root {
		return 0
	}
	//左右递归找最大路径和
	leftMax := getMax(outOrderHelper(root.Left), 0)
	rightMax := getMax(outOrderHelper(root.Right), 0)

	res = getMax(res, root.Val+leftMax+rightMax)
	fmt.Println("root.val=", root.Val, res)
	return getMax(root.Val+getMax(leftMax, rightMax), 0)
}

//方法2：内置函数法
func maxPathSum2(root *TreeNode) int {
	res := math.MinInt32

	//内置函数，可以直接使用 res，而不用考虑复杂的递归规程中的传参
	var outOrderHelper func(root *TreeNode) int
	outOrderHelper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getMax(outOrderHelper(root.Left), 0)
		right := getMax(outOrderHelper(root.Right), 0)

		//更新res
		res = getMax(res, left+root.Val+right)
		//返回当前结点的最大路径和
		return getMax(root.Val+getMax(left, right), 0)
	}

	outOrderHelper(root)
	return res
}

//方法3：指针法
func maxPathSum3(root *TreeNode) int {
	res := math.MinInt32
	outOrderHelper3(root, &res)
	return res
}

func outOrderHelper3(root *TreeNode, res *int) int {
	if nil == root {
		return 0
	}
	//左右递归找最大路径和
	leftMax := getMax(outOrderHelper3(root.Left, res), 0)
	rightMax := getMax(outOrderHelper3(root.Right, res), 0)

	*res = getMax(*res, root.Val+leftMax+rightMax)
	fmt.Println("root.val=", root.Val, res)
	return getMax(root.Val+getMax(leftMax, rightMax), 0)
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
