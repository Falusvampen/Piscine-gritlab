package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
} */

func ListAt(l *NodeL, pos int) *NodeL {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}
	if pos <= index {
		for l != nil {
			if count == pos {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
