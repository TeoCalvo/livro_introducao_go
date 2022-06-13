package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {

		txt := ""
		if divideBy(i, 3) {
			txt += "Fizz"
		}
		if divideBy(i, 5) {
			txt += "Buzz"
		}
		if txt == "" {
			txt = strconv.Itoa(i)
		}

		fmt.Println(txt)

	}

}

func divideBy(number, divider int) bool {
	return number%divider == 0
}
