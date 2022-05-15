package piscine

func UltimateDivMdsod(a *int, b *int) {

	nice := *a
	*a = *a / *b
	*b = nice % *b
}
