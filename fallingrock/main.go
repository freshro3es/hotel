package main

import (
	"fmt"
)

func main() {
	for i := 50; i <= 100; i++ {
		fmt.Println(i, ": ", 5000/i+i-1)
	}
}
