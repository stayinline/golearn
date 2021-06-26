package main

var maps map[int]int // k:路径和，v:路径和出现的次数

func pathSum(root *TreeNode, targetSum int) int {
	maps = make(map[int]int)
	maps[0] = 1
	return pathSumHelper(root, targetSum, 0)
}

func pathSumHelper(root *TreeNode, targetSum, curSum int) int {
	if nil == root {
		return 0
	}
	ret := 0
	curSum += root.Val                                  //更新当前和
	ret += maps[curSum-targetSum]                       //更新当前节点对应的路径和出现的次数
	maps[curSum] = maps[curSum] + 1                     //将当前节点路径和及其出现的次数存入map
	ret += pathSumHelper(root.Left, targetSum, curSum)  //递归左树
	ret += pathSumHelper(root.Right, targetSum, curSum) //递归右树
	maps[curSum] = maps[curSum] - 1                     //这里很关键，回溯的思想，从路径和中撤销当前结点
	return ret
}

func main() {

}
