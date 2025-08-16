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
	sideLength float64
}

func main() {
	shapes := []shape{
		triangle{base: 3, height: 4},
		square{sideLength: 5},
	}

	for _, s := range shapes {
		fmt.Println(s.getArea())
	}
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
