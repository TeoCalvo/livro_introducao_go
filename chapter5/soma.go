package main

import "fmt"

func main() {

	var notas [5]float64
	notas[0] = 98
	notas[1] = 93
	notas[2] = 77
	notas[3] = 82
	notas[4] = 83

	fmt.Println("Média:", soma(notas[:])/float64(len(notas)))

}

func soma(x []float64) float64 {
	var total float64
	for _, value := range x {
		total += value
	}
	return total
}
