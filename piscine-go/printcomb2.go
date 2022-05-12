package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for gomb := '0'; gomb <= '9'; gomb++ {
		for lasagna := '0'; lasagna <= '9'; lasagna++ {
			for touch := '0'; touch <= '9'; touch++ {
				for kebab := '0'; kebab <= '9'; kebab++ {
					if gomb < lasagna {
						if lasagna < touch {
							if touch < kebab {
								z01.PrintRune(gomb)
								z01.PrintRune(lasagna)
								z01.PrintRune(' ')
								z01.PrintRune(touch)
								z01.PrintRune(kebab)
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}
