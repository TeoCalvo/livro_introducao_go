package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		res := i
		i += 2
		return res
	}
}

func main() {
	nexOdd := makeOddGenerator()

	fmt.Println(nexOdd())
	fmt.Println(nexOdd())
	fmt.Println(nexOdd())
	fmt.Println(nexOdd())

}
