package fndcomdiv

// Функция на вход получает массив положительных целых чисел произвольной длины, например [42, 12, 18].
// На выходе возвращает массив чисел, которые являются общими делителями для всех указанных числе. В примере это будет [1, 2, 3, 6].
func FindCommonDivisors(numbers []int) []int {
	divisors := make(map[int]int)

	for _, number := range numbers {
		for i := 1; i*i <= number; i++ {
			if number%i == 0 {
				divisors[i]++
				// Если i не является квадратным корнем числа, добавляем второй делитель.
				if i*i != number {
					divisors[number/i]++
				}
			}
		}
	}

	var commonDivisors []int
	for divisor, count := range divisors {
		if count == len(numbers) {
			commonDivisors = append(commonDivisors, divisor)
		}
	}
	return commonDivisors
}
