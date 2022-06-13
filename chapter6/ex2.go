package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(half(x))

}

func half(x int) (int, bool) {

	res := x / 2
	even := x%2 == 0
	return res, even

}
