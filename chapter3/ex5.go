package main

import "fmt"

func main() {
	tempFahrenheit := 0.

	fmt.Print("Entre com um número em Fahrenheit:")
	fmt.Scanf("%f", &tempFahrenheit)

	tempCelsius := (tempFahrenheit - 32) * 5.0 / 9.0
	fmt.Println(tempCelsius, "˚C")
}
