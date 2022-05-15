package piscine

func StrRev(s string) string {
	strang := []rune(s)
	// array1 = 0     array2 är längd av strängen - 1      om array1 är mindre än längden av strängen -1    +1 på array1   och -1 på array2
	for array1, array2 := 0, len(strang); array1 < array2; array1, array2 = array1+1, array2-1 {
		strang[array1], strang[array2] = strang[array2], strang[array1]
	}
	return string(strang)
}
