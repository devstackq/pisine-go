package piscine

func ListClear(l *List) {

	for l.Head != nil {
		tmp := l.Head

		l.Head = nil
		l.Tail = nil
		l.Head.Next = tmp
	}

}

//tmp - tail - head
// }
// node := l.Tail
// node.Next = nil
// 