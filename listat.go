package piscine

func ListAt(l *NodeL, pos int) *NodeL {

	count := 0

	for l != nil {
		count++
		if pos == count-1 {
			return l // return l, value
		}
		l = l.Next // start next func
	}
	return nil
}
