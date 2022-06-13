package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Retangle struct {
	x1, x2, y1, y2 float64
}

func (r *Retangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	circle := Circle{x: 0, y: 0, r: 10}
	retangle := Retangle{x1: 0, x2: 10, y1: 0, y2: 10}

	fmt.Println("Área do círculo de raio 10:", circle.area())
	fmt.Println("Área do retangulo:", retangle.area())
}
