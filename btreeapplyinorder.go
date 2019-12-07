package piscine

// 4,1,7,5

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	//
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		// recursion call, left side, left side - always low
		//if func view, tree, going to left, then root, then right, before all node check recursion call
		// and return order result
		//left, root, right
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}
