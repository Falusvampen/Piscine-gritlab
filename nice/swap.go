package piscine

func Swap(a *int, b *int) {
	var nice int = *a
	var dude int = *b
	*a = dude
	*b = nice
}
