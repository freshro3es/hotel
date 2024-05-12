package printtab

import (
	"fmt"
	"strconv"
)

func calcLen(n int) int {
	return len(strconv.Itoa(n))
}

func makeSpaces(distance int) string {
	spaces := ""
	for i := 0; i < distance; i++ {
		spaces += " "
	}
	return spaces
}

func printNumber(number int, width int) string {
	return fmt.Sprintf("%*d", width, number)
}

func PrintTable(n int) string {
	var result string
	for y := 0; y <= n; y++ {
		for x := 0; x <= n; x++ {
			if y == 0 && x == 0 {
				result += makeSpaces(calcLen(n))
				continue
			}
			if y == 0 {
				result += printNumber(x, calcLen(n*x)+1)
				continue
			}
			if x == 0 {
				result += printNumber(y, calcLen(n))
				continue
			}
			result += printNumber(x*y, calcLen(n*x)+1)
		}
		result += "\n"
	}
	fmt.Print(result)
	return result
}
