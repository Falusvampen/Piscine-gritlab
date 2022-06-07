package piscine

func Map(f func(int) bool, a []int) []bool {
	nice := make([]bool, len(a))
	for i, e := range a {
		nice[i] = f(e)
	}
	return nice
}
