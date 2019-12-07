package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

	head := l.Head

	for head != nil {
		if head == ref {
			comp(ref, head)
		}
		head = head.Next
	}
	return &ref
}
