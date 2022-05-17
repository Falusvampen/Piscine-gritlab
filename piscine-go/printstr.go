package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	nice := []rune(s)
	for hello := 0; hello < len(s); hello++ {
		z01.PrintRune(rune(nice[hello]))
	}
}
