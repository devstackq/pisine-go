package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	//1 Height - BtreeCount
	//2call printLev
	level := BTreeLevelCount(root)

	if root == nil {
		return
	}

	for i := 0; i < level; i++ {
		printLevel(root, f, i, 0)
	}
}

//1 call i 0, level 0,
// 2 call i1, count 0, return root Data
//3call left count +1, iteration 1, print rootLeftData -1

func printLevel(root *TreeNode, f func(...interface{}) (int, error), level, count int) {
	if level == count {
		f(root.Data)

	}
	if root.Left != nil {
		printLevel(root.Left, f, level, count+1)
	}
	if root.Right != nil {
		printLevel(root.Right, f, level, count+1)
	}

}
