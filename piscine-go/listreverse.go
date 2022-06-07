package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
}
type List struct {
 Head *NodeL
 Tail *NodeL
} */

func ListReverse(l *List) {
	current := l.Head
	prev := l.Tail
	prev = nil
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
