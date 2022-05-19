package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args[0]
	z := []rune(x)
	for a, b := range z {
		if a > 1 {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}
