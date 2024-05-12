package calcdec

import (
	"strconv"
)

// Принимает целое число в качестве аргумента и возвращает строку, содержащую это число и слово "компьютер" в нужном склонении по падежам в зависимости от числа.
func CalcDeclination(n int) string {
	nStr := strconv.Itoa(n)
	if (n%100) > 10 && (n%100) < 15 {
		return nStr + " компьютеров"
	}
	switch n % 10 {
	case 1:
		return nStr + " компьютер"
	case 2, 3, 4:
		return nStr + " компьютера"
	case 0, 5, 6, 7, 8, 9:
		return nStr + " компьютеров"
	}
	return ""
}
