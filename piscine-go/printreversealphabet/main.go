package main

import "github.com/01-edu/z01"

func main() {
	var alphabet string = "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}
