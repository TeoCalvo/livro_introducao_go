package main

import (
	"fmt"

	"github.com/teocalvo/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println("3 + 5 = ", total)
}
