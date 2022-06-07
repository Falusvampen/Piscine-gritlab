package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
}
type List struct {
 Head *NodeL
 Tail *NodeL
} */

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(List *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := List.Head
	for iterator != nil {
		if comp(iterator.Data, ref) {
			return &iterator.Data
		}

		iterator = iterator.Next
	}
	return nil
}
