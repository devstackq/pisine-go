package piscine

// 4,1,7,5

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {

	if root != nil {
		BTreeApplyPostorder(root.Left, f)
		// recursion call, left side, left side - always low
		//if func view, tree, going to left, then root, then right, before all node check recursion call
		// and return order result
		//left, root, right
		BTreeApplyPostorder(root.Right, f)
		f(root.Data)
	}
	//same inorder, but, sequence another -> left, right, root recursion -> by tree
}
