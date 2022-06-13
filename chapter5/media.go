package main

import "fmt"

func main() {
	notas := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	fmt.Println("Média:", media(notas[:]))
}

func soma(x []float64) float64 {
	var total float64
	for _, v := range x {
		total += v
	}
	return total
}

func media(x []float64) float64 {
	total := soma(x)
	n := float64(len(x))
	return total / n
}
