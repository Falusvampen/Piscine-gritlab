package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
}
type List struct {
 Head *NodeL
 Tail *NodeL
} */

func ListRemoveIf(l *List, data_ref interface{}) {
	head := l.Head
	tail := l.Head

	for tail != nil && tail.Data == data_ref {
		l.Head = tail.Next
		tail = l.Head
	}
	for tail != nil {
		for tail != nil && tail.Data != data_ref {
			head = tail
			tail = tail.Next
		}

		if tail == nil {
			return
		}
		head.Next = tail.Next
		tail = head.Next
	}
}
