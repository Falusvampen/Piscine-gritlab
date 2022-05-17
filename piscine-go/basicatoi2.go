package piscine

func BasicAtoi2(s string) int {
	str := 0
	// leaves the index value and declares nice as s
	for _, nice := range s {
		// if the nice is larger than 47 and smaller than 58
		if int(rune(nice)) > 47 && int(rune(nice)) < 58 {
			// declare a = 0 and a simple for loop
			a := 0
			for i := '1'; i <= nice; i++ {
				a++
			}
			//
			str = str*10 + a
		} else {
			return 0
		}
	}
	return str
}
