package piscine

func Any(f func(string) bool, a []string) bool {
	result := false
	nice := make([]bool, len(a))
	for i, e := range a {
		nice[i] = f(e)
		if nice[i] == true {
			result = true
		}
	}
	return result
}
