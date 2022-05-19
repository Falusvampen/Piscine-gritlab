package piscine

func Swap(a *int, b *int) {
	nice := *a
	dude := *b
	*a = dude
	*b = nice
}
