package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

func writeArea(f form) {
	fmt.Printf("The area is %0.2f \n", f.area())
}

func (r ret) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type ret struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

func main() {
	r := ret{10, 15}
	writeArea(r)

	c := circle{10}
	writeArea(c)
}
