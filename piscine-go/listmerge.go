package piscine

/* type NodeL struct {
 Data interface{}
 Next *NodeL
}
type List struct {
 Head *NodeL
 Tail *NodeL
} */

func ListMerge(List1 *List, List2 *List) {
	if List2 == nil || List1 == nil {
		return
	}
	if List1.Head == nil {
		List1.Tail = List2.Head
		List1.Head = List2.Head
		return
	}
	List1.Tail.Next = List2.Head
}
