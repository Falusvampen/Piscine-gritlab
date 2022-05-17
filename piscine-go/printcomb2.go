package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for gomb := '0'; gomb <= '9'; gomb++ {
		for kebab := '0'; kebab <= '9'; kebab++ {
			for sushi := '0'; sushi <= '9'; sushi++ {
				for taco := '0'; taco <= '9'; taco++ {
					z01.PrintRune(gomb)
					z01.PrintRune(kebab)
					z01.PrintRune(' ')
					if gomb >= sushi {
						sushi = gomb
						if kebab >= taco {
							taco = kebab + 1
						}
					}
					if taco > '9' {
						taco = '0'
						sushi++
					}
					z01.PrintRune(sushi)
					z01.PrintRune(taco)
					if gomb == '9' && kebab == '8' {
						z01.PrintRune('\n')
						return
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
