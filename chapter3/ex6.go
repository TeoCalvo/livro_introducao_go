package main

import "fmt"

func main() {
	pes := 0.
	fmt.Print("Entre com um valor em pés: ")
	fmt.Scanf("%f", &pes)

	metros := pes * 0.3048
	fmt.Println(metros, "metros")
}
