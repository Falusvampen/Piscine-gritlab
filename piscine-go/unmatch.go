package piscine

func Unmatch(a []int) int {
	var count int
	for _, e := range a {
		count = 0
		for _, e2 := range a {
			if e == e2 {
				count++
			}
		}
		if count%2 != 0 {
			return e
		}
	}
	return -1
}
