package piscine

func printOrderTraversal(root *TreeNode, x int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if x == 1 {
		f(root.Data)
	} else if x > 1 {
		printOrderTraversal(root.Left, x-1, f)
		printOrderTraversal(root.Right, x-1, f)
	}
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	count := BTreeLevelCount(root)
	for i := 1; i <= count; i++ {
		printOrderTraversal(root, i, f)
	}
}
