package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	b := 1
	if n < 0 {
		b = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		c := (n / 10) * b
		if c != 0 {
			PrintNbr(c)
		}
		bork := (n % 10 * b) + '0'
		z01.PrintRune(rune(bork))
	} else {
		z01.PrintRune('0')
	}
}

func PrintStr(s string) {
	for i := range s {
		z01.PrintRune(rune(s[i]))
	}
}

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := new(point)

	setPoint(points)

	PrintStr("x = ")
	PrintNbr(points.x)
	PrintStr(", ")
	PrintStr("y = ")
	PrintNbr(points.y)
	PrintStr("\n")
}
