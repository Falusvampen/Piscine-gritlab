package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for gomb := '0'; gomb <= '9'; gomb++ {
		for lasagna := '0'; lasagna <= '9'; lasagna++ {
			for touch := '0'; touch <= '9'; touch++ {
				if gomb < lasagna {
					if lasagna < touch {
						z01.PrintRune(gomb)
						z01.PrintRune(lasagna)
						z01.PrintRune(touch)
						if gomb != '7' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						} else {
							z01.PrintRune('\n')
						}
					}
				}
			}
		}
	}
}
