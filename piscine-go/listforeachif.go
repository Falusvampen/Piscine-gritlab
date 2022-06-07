package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
}
type List struct {
 Head *NodeL
 Tail *NodeL
} */

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	x := l.Head
	for x != nil {
		if cond(x) {
			f(x)
		}
		x = x.Next
	}
}
