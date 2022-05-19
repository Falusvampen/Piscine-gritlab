package piscine

func UltimateDivMod(a *int, b *int) {
	var x int = *a
	var y int = *b
	*a = x / y
	*b = x % y
}
