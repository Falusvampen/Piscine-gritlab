package piscine

/* type NodeI struct {
 Data int
 Next *NodeI
} */

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	n2 = ListSort(n2)
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}
	if n1.Data > n2.Data {
		n2.Next = SortedListMerge(n1, n2.Next)
		return n2
	} else {
		n1.Next = SortedListMerge(n1.Next, n2)
		return n1
	}
}
