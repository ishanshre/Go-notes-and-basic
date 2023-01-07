package main

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circum() float64 {
	return s.length * 4
}
