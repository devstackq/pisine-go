package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}

	lcond := BTreeIsBinary(root.Left)
	rcond := BTreeIsBinary(root.Right)

	if lcond == true && rcond == true {
		return true
	}

	return false
}
