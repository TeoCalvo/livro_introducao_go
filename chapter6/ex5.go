package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("fib(%d) = %d\n", n, fib(n))

}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
