package piscine

func ListLast(l *List) interface{} {
	node := l.Head

	if node == nil {
		return node
	}
	for node.Next != nil {
		node = node.Next
	}
	return node.Data

	// if l.Tail != nil {
	// 	return l.Tail.Data
	// }
	// return l.Tail

}
