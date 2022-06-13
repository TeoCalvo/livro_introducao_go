package main

import "fmt"

func main() {

	for i := 1; i <= 11; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		case 4:
			fmt.Println("Quatro")
		case 5:
			fmt.Println("Cinco")
		case 6:
			fmt.Println("Seis")
		case 7:
			fmt.Println("Sete")
		case 8:
			fmt.Println("Oito")
		case 9:
			fmt.Println("Nove")
		case 10:
			fmt.Println("Dez")
		default:
			fmt.Println("Número desconhecido!")
		}
	}
}
