package printtab

import (
	"testing"
)

func TestMuliplTable(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected string
	}{
		{
			name: "Example 1",
			n:    5,
			expected: `  1  2  3  4  5
1 1  2  3  4  5
2 2  4  6  8 10
3 3  6  9 12 15
4 4  8 12 16 20
5 5 10 15 20 25
`,
		},
		{
			name: "Example 2",
			n:    3,
			expected: `  1 2 3
1 1 2 3
2 2 4 6
3 3 6 9
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PrintTable(test.n)
			if result != test.expected {
				t.Errorf("Test case %s failed: expected\n%s\nbut got\n%s", test.name, test.expected, result)
			}
		})
	}
}

// Пример использования функции MultiplTable.
func Example_multiplTable() {
	// Выведем таблицу умножения 5х5
	PrintTable(5)
	// Output:
	//   1  2  3  4  5
	// 1 1  2  3  4  5
	// 2 2  4  6  8 10
	// 3 3  6  9 12 15
	// 4 4  8 12 16 20
	// 5 5 10 15 20 25
}
