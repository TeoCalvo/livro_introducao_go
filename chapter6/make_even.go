package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		res := i
		i += 2
		return res
	}
}

func main() {

	nextEven := makeEvenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

}
