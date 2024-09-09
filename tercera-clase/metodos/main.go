package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radio
}

type Circle struct {
	Radio float64
}

type Square struct {
	Lado float64
}

func (s Square) Area() float64 {
	return s.Lado * s.Lado
}

func (s Square) Perimeter() float64 {
	return s.Lado * 4
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c *Circle) setRadio(radio float64) {
	c.Radio = radio
}

type Number int

func (n Number) IsPositive() bool {
	return n > 0
}

func ImprimirArea(c Shape) {
	fmt.Println("Area", c.Area())
}

func main() {
	circulo := Circle{Radio: 5}
	cuadrado := Square{Lado: 5}
	n1 := Number(10)

	ImprimirArea(circulo)
	ImprimirArea(cuadrado)

	fmt.Println(circulo.Area())
	fmt.Println(n1.IsPositive())

	circulo.setRadio(10)
	fmt.Println(circulo.Radio)
}
