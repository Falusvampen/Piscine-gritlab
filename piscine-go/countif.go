package piscine

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	nice := make([]bool, len(tab))
	for i, e := range tab {
		nice[i] = f(e)
		if nice[i] == true {
			result++
		}
	}
	return result
}
