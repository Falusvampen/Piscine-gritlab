package piscine

func ListSize(l *List) int {
	x := l.Head
	count := 0
	for x != nil {
		count++
		x = x.Next
	}
	return count
}
