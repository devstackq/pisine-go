package piscine

func ListSize(l *List) int {

	size := 0

	for l.Head != nil {
		size++
		l.Head = l.Head.Next
	}
	return size
}

// last := l.Head
// for {
// 	if last == nil || last.Next == nil {
// 		break
// 	}
// 	last = last.Next
// 	size++
