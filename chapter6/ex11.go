package main

import "fmt"

func main() {

	x, y := 1, 2
	swap(&x, &y)
	fmt.Printf("x=%d e y=%d\n", x, y)

}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
