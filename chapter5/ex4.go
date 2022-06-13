package main

import "fmt"

func main() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println("Menor:", menor(x))

}

func menor(x []int) int {
	n := x[0]

	for _, v := range x[1:] {
		if v < n {
			n = v
		}
	}
	return n
}
