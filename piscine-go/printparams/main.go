package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, e := range arguments {
		for i := 0; i < len([]rune(e)); i++ {
			z01.PrintRune(rune(e[i]))
		}
		z01.PrintRune('\n')
	}
}
