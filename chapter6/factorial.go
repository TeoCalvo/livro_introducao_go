package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(x int) int {
	if x == 1 {
		return x
	} else {
		return x * factorial(x-1)
	}
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	f := factorial(n)

	fmt.Printf("%d! = %d\n", n, f)

}
