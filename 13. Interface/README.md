# GO Interfaces
- An interface types is a custom type that are the collection of sets of method signatures
- It is an abstract type and does not have any implementation
- It describes the expected behaviour of a type
- Interface in go can also be described as duck typing
- "if it walks like a duck and quack like a duck, then it must be a duck"
- In code below, I have create two struct square and circle with two different receiver function with same name. So if we don't use interface then we need to call them individually. Its not dynamic.
- So we use interface make our code small and clean.
- Most important thing is if square does not have one of the method then interface will not work



```
package main

import (
    "math"
    "fmt"
    )

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) circum() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circum() float64 {
	return s.length * 4
}

type shape interface {
	area() float64
	circum() float64
}

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f\n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f\n", s, s.circum())
}

func main() {
	printShapeInfo(square{length: 10})
	shapes := []shape{
		square{length: 2},
		square{length: 3.3},
		circle{radius: 4.4},
		circle{radius: 3},
	}
	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("_____")
	}
}
```