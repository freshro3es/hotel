package fndcomdiv

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestFindCommonDivisors(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected []int
	}{
		{
			name:     "Example 1",
			numbers:  []int{42, 12, 18},
			expected: []int{1, 2, 3, 6},
		},
		{
			name:     "Example 2",
			numbers:  []int{30, 15, 25},
			expected: []int{1, 5},
		},
		{
			name:     "Example 3",
			numbers:  []int{13, 17, 19},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindCommonDivisors(test.numbers)
			sort.Ints(result)
			sort.Ints(test.expected)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test case %s failed: expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

// Пример использования функции FindCommonDivisors.
func Example_findCommonDivisors() {
	// Вызов функции для массива [42, 12, 18].
	result := FindCommonDivisors([]int{42, 12, 18})
	sort.Ints(result)
	fmt.Println(result)
	// Output: [1 2 3 6]
}
