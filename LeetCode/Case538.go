package main

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	outOrderHelper2(root)
	return root
}

func outOrderHelper2(root *TreeNode) {
	if nil == root {
		return
	}
	outOrderHelper2(root.Right)
	sum += root.Val
	root.Val = sum
	outOrderHelper2(root.Left)
}

func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	var outOrderHelper func(*TreeNode)
	outOrderHelper = func(node *TreeNode) {
		if node != nil {
			outOrderHelper(node.Right)
			sum += node.Val
			node.Val = sum
			outOrderHelper(node.Left)
		}
	}
	outOrderHelper(root)
	return root
}

func main() {

}
