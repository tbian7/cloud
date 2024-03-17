package lc

import "strings"

func intToRoman(num int) string {
	var sb strings.Builder

	sb.WriteString(digit(num/1000, 'M', 0, 0))
	num = num % 1000
	sb.WriteString(digit(num/100, 'C', 'D', 'M'))
	num = num % 100
	sb.WriteString(digit(num/10, 'X', 'L', 'C'))
	num = num % 10
	sb.WriteString(digit(num, 'I', 'V', 'X'))
	return sb.String()
}

func digit(d int, one, five, ten rune) string {
	var sb strings.Builder
	switch d {
	case 0, 1, 2, 3:
		for i := 0; i < d; i++ {
			sb.WriteRune(one)
		}
	case 4:
		sb.WriteRune(one)
		sb.WriteRune(five)
	case 5, 6, 7, 8:
		sb.WriteRune(five)
		for i := 5; i < d; i++ {
			sb.WriteRune(one)
		}
	case 9:
		sb.WriteRune(one)
		sb.WriteRune(ten)
	}
	return sb.String()
}
