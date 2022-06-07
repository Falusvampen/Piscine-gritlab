package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	SortIntegerTable(table)
	return table[2]
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[j] > table[i] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
