package prmnum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrimeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		end      int
		expected []int
	}{
		{
			name:     "Example 1",
			start:    11,
			end:      20,
			expected: []int{11, 13, 17, 19},
		},
		{
			name:     "Example 2",
			start:    20,
			end:      30,
			expected: []int{23, 29},
		},
		{
			name:     "Example 3",
			start:    1,
			end:      10,
			expected: []int{2, 3, 5, 7},
		},
		// Добавьте другие тестовые случаи по необходимости
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PrimeNumbers(test.start, test.end)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test case %s failed: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

// Пример использования функции PrimeNumbers.
func Example_primeNumbers() {
	// На вход переданы 2 числа: от 11 до 20.
	fmt.Println(PrimeNumbers(11, 20))
	// Output: [11 13 17 19]
}
