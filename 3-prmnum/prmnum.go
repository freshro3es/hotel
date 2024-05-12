package prmnum

func isPrime(n int) bool {
	if n%2 == 0 && n != 2 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Функция возвращает массив простых чисел в диапазоне (2 числа - минимальное и максимальное) заданных чисел.
// Диапазон считается включая граничные значения.
func PrimeNumbers(start int, end int) []int {
	if end < 2 || start > end {
		return []int{}
	}
	if start < 2 {
		start = 2
	}

	result := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}
	return result
}
