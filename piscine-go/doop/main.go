package main

import "os"

func Atoi(number string) (int, bool) {
	result := 0
	neg := false
	if len(number) == 0 {
		return 0, false
	}
	if number[0] == '+' {
		number = number[1:]
	} else if number[0] == '-' {
		neg = true
		number = number[1:]
	}
	for _, e := range number {
		if e >= '0' && e <= '9' {
			result *= 10
			result += int(e) - 48
		} else {
			return 0, false
		}
	}
	if neg {
		result *= -1
	}
	return result, true
}

func Itoa(number int) string {
	result := ""
	neg := false
	if number == 0 {
		return "0"
	}
	if number < 0 {
		number *= -1
		neg = true
	}
	for number > 0 {
		result = string(number%10+48) + result
		number /= 10
	}
	if neg {
		result = "-" + result
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	a, r := Atoi(args[0])
	b, er := Atoi(args[2])
	if !r || !er {
		return
	}
	switch args[1] {
	case "*":
		result := a * b
		if a == 0 || (result/a) == b {
			os.Stdout.WriteString(Itoa(result))
			os.Stdout.WriteString("\n")
		}
	case "+":
		result := a + b
		if (result > a) == (b > 0) {
			os.Stdout.WriteString(Itoa(result))
			os.Stdout.WriteString("\n")
		}
	case "-":
		result := a - b
		if (result < a) == (b > 0) {
			os.Stdout.WriteString(Itoa(result))
			os.Stdout.WriteString("\n")
		}
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0")
			os.Stdout.WriteString("\n")
		} else {
			result := a / b
			os.Stdout.WriteString(Itoa(result))
			os.Stdout.WriteString("\n")
		}
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0")
			os.Stdout.WriteString("\n")
		} else {
			os.Stdout.WriteString(Itoa(a % b))
			os.Stdout.WriteString("\n")
		}
	default:
		return
	}
}
