package main

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	diameterOfBinaryTreeHelper(root, &res)
	return res
}

func diameterOfBinaryTreeHelper(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := diameterOfBinaryTreeHelper(root.Left, res)
	right := diameterOfBinaryTreeHelper(root.Right, res)
	*res = getMax(*res, left+right)
	return getMax(left, right) + 1
}

func main() {
	diameterOfBinaryTree(nil)

}
