package main

import "math"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) circum() float64 {
	return 2 * math.Pi * c.radius
}
