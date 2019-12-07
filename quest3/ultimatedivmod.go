package piscine

func UltimateDivMod(a *int, b *int) {
	n := *a
	*a = n / *b
	*b = n % *b
}
