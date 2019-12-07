package piscine

func ListSize(l *List) int {

	size := 1
	last := l.Head
	for {
		if last == nil || last.Next == nil {
			break
		}
		last = last.Next
		size++
	}
	return size

}

// count := 0
// for l.Head != nil {

// 	count++
// 	l.Head = l.Head.Next

// }
// return count

// -
//  func ListSize(l *List) int {
