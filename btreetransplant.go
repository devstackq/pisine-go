package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node.Parent == nil {
		return rplc
	}
	if node.Parent.Left == node {
		rplc.Parent = node.Parent
		node.Parent.Left = rplc
	}
	if node.Parent.Right == node {
		rplc.Parent = node.Parent
		node.Parent.Right = rplc
	}
	return root
}
