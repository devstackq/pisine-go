package piscine

func ListRemoveIf(l *List, data_ref interface{}) {

	tmp := l.Head  // write head - tmp var
	prev := l.Head // 2 head - prev var

	// for loop work, head != nil and  tmp.Data == data_ref
	for tmp != nil && tmp.Data == data_ref {
		l.Head = tmp.Next //Head change point - NExt
		tmp = l.Head      // tmp
	}

	for tmp != nil {
		for tmp != nil && tmp.Data != data_ref {
			prev = tmp
			tmp = tmp.Next
		}

		if tmp == nil {
			return
		}

		prev.Next = tmp.Next
		tmp = prev.Next
	}
}
