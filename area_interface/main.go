package main

import "fmt"

type shape interface {
	printArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLenght float64
}

func main() {

	t := triangle{
		base:   10.0,
		height: 5.0,
	}

	fmt.Println("Area of the triangle is:", t.printArea())
	sq := square{
		sideLenght: 10.0,
	}

	fmt.Println("Area of the square is :", sq.printArea())
}

func (sq square) printArea() float64 {
	return sq.sideLenght * sq.sideLenght
}

func (t triangle) printArea() float64 {
	return 0.5 * t.base * t.height
}
