package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	x := stringToInt(os.Args[1:]...)

	fmt.Println(maior(x...))

}

func maior(x ...int) int {
	m := x[0]
	for _, v := range x[1:] {
		if m < v {
			m = v
		}
	}
	return m
}

func stringToInt(x ...string) []int {

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
