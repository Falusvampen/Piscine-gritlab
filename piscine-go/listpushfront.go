package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		newNode := &NodeL{Data: data}
		newNode.Next, l.Head = l.Head, newNode
	}
}
