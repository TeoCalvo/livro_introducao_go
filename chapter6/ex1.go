package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1:]

	x := convSrt(n...)
	fmt.Println(soma(x...))
}

func convSrt(x ...string) []int {
	y := []int{}
	for _, v := range x {
		new_v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		y = append(y, new_v)
	}
	return y
}

func soma(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
