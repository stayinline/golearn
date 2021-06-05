package main

import (
	"fmt"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

func main() {

	fmt.Println(math.MinInt64, math.MaxInt64)
	fmt.Println(math.MinInt32, math.MaxInt32)

}
