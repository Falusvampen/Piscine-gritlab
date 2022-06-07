package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, word := range table {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
