package piscine

func ListLast(l *List) interface{} {
	node := l.Head

	//head
	//nil true
	if node == nil {
		return nil
	}
	// for , poka != nil
	for node.Next != nil {
		node = node.Next // write node - last node.Next
	}
	return node.Data // return last data

	// if l.Tail != nil {
	// 	return l.Tail.Data
	// }
	// return l.Tail
}
