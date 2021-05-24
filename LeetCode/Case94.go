package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (res []int) {
	inorderHelper(root, &res)
	return res
}

func inorderHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderHelper(root.Right, res)
}

func inorderTraversal1(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

func main() {

}
