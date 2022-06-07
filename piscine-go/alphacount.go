package piscine

func AlphaCount(s string) int {
	var count int
	for _, j := range s {
		if j >= 'a' && j <= 'z' || j >= 'A' && j <= 'Z' {
			count++
		}
	}
	return count
}
