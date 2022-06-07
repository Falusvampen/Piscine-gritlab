package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := range s {
		z01.PrintRune(rune(s[i]))
	}
}
