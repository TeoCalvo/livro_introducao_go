package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Olá, meu nome é", p.Name)
}

type Andriod struct {
	Person
	Model string
}

func main() {
	a := new(Andriod)
	a.Name = "TéoBot"
	a.Model = "Fodase"
	a.Talk()
}
