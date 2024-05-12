package calcdec

import (
	"fmt"
	"testing"
)

func TestCalcDeclination(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "1 компьютер"},
		{2, "2 компьютера"},
		{3, "3 компьютера"},
		{4, "4 компьютера"},
		{5, "5 компьютеров"},
		{11, "11 компьютеров"},
		{12, "12 компьютеров"},
		{13, "13 компьютеров"},
		{14, "14 компьютеров"},
		{20, "20 компьютеров"},
		{21, "21 компьютер"},
		{22, "22 компьютера"},
		{25, "25 компьютеров"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := CalcDeclination(tc.input)
			if result != tc.expected {
				t.Errorf("For input %d, expected %q, but got %q", tc.input, tc.expected, result)
			}
		})
	}
}

// Пример использования функции СalcDeclination.
func Example_calcDeclination() {
	// Вызов функции для числа 5.
	fmt.Println(CalcDeclination(5))
	// Output: 5 компьютеров
}
