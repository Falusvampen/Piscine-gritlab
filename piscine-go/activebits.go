package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
