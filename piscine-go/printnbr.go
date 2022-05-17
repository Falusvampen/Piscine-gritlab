/*package piscine

import "github.com/01-edu/z01"

func PrintNbr(a int) {
	//i detta exempel använder vi ett negativt tal -123
	//första check,
	//defineras variabeln b som 1
	b := 1
	//om a är mindre än noll så är det ett negativt tal vilket leder till att b ändrar värde till negativt ett (-1)
	if a < 0 {
		//denna negativa etta är endast till för att användas i multiplikation för att ändra värde på a, tex -123(-1) = 123, ett positivt tal
		b = -1
		//i det här fallet printas även "-"
		z01.PrintRune('-')
	}
	//nästa check
	//eftersom vårt exempel -123 inte är lika med 0 så exekveras denna kod
	if a != 0 {
		//definera ny variabel c vilket i det här fallet är lika med (-123/10)*-1    (minus ett är lika med variabeln b)
		// -123/10 = -12,3,        -1(-12,3)=12,3         alltså variabel c är i detta fall = 12 eftersom decimaler kapas
		c := (a / 10) * b
		//om 12 är inte lika med 0 vilket betyder att 12(alltså c) printas genom huvudfunktionen
		if c != 0 {
			PrintNbr(c)
		}
		//här defineras variabeln d som -123 % 10 = -3 * b(vilket är -1)    -3(-1) = 3
		//(a % 10 * b) = 3           3+0=3
		d := (a % 10 * b) + '0'
		//printar då d vilket vi löste ut som 3
		z01.PrintRune(rune(d))

		//denna kod exekveras endast ifall de två checkarna innan ej stämde, alltså printas inget just nu då koden inte är lika med 0
	} else {
		z01.PrintRune('0')
	}
	//i hela denna funktion printar alltså check ett "-"      check två     12 och 3 printas
}
*/
package piscine

import "github.com/01-edu/z01"

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
