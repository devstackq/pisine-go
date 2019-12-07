package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	// if root == nil {
	// 	return nil
	// }

	// if 4 < 7 search recursion elem - Left side
	// else if serach recursion - right side, and return him data

	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	} else if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	} else {
		return root
	}
}
