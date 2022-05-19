package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
	width  float64
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

func (s square) getArea() float64 {
	return (s.length * s.width)
}

func printArea(s shape) {
	fmt.Printf("%g\n", s.getArea())
}

func main() {
	t := triangle{2, 3}
	s := square{2, 3}
	printArea(t)
	printArea(s)
}