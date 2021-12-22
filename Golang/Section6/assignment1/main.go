package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 4}
	printArea(sq)

	tr := triangle{
		base: 4,
		height: 10,
	}
	printArea(tr)
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println("Area: ", area)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}