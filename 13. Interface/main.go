package main

import "fmt"

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
