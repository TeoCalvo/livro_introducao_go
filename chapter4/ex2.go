package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if divideBy(i, 3) {
			fmt.Println(i)
		}
	}
}

func divideBy(number, divider int) bool {
	return number%divider == 0
}
