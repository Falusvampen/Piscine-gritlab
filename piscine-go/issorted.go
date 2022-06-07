package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	val1 := 1
	val2 := 1
	val3 := 1
	for x, y := range tab {
		if x != len(tab)-1 {
			if f(y, tab[x+1]) < 0 {
				val1++
			}
			if f(y, tab[x+1]) > 0 {
				val2++
			}
			if f(y, tab[x+1]) == 0 {
				val3++
			}
		}
	}
	return val1 == len(tab) || val2 == len(tab) || val3 == len(tab)
}
