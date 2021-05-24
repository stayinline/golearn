package main




func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if nil == root1 {
		return root2
	} else if nil == root2 {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

func main() {

}
