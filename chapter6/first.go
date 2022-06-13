package main

import "fmt"

func average(xs []float64) float64 {
	panic("Não implementada!")
}

func main() {

	var notas [5]float64
	notas[0] = 98
	notas[1] = 93
	notas[2] = 77
	notas[3] = 82
	notas[4] = 83

	fmt.Println("Media:", average(notas[:]))

}
