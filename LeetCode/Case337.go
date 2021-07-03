package main

func rob(root *TreeNode) int {
	val := dfsHelper(root)
	return getMax(val[0], val[1])
}

//返回值数组：
// 第一个是元素是选择当前结点的结果
// 第二个元素是不选择当前结点的结果
func dfsHelper(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	l := dfsHelper(cur.Left)
	r := dfsHelper(cur.Right)

	//选择当前节点：将当前结点的值加入
	selected := cur.Val + l[1] + r[1]

	//不选择当前结点：那么就是左孩子和右孩子中的最大值
	// 左孩子最大值：就是返回值中的最大值，右孩子同理
	notSelected := getMax(l[0], l[1]) + getMax(r[0], r[1])
	return []int{selected, notSelected}
}

func main() {

}
