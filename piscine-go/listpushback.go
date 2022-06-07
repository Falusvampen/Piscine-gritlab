package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		return
	}
	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}
