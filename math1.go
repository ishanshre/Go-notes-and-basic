package main

import (
	"fmt"
	"math"
)

func main() {
	// ceil  ---> give the ceiling the value
	a := 2.000001
	fmt.Println(a)
	fmt.Println(math.Ceil(a))
	// min ---> smaller one between,, max ----> should give the larger between two
	fmt.Println(math.Floor(3.9)) // floor is oposite of the ceil
	fmt.Println(math.Min(2, 3.9))
	fmt.Println(math.Max(3, 4))
	fmt.Println(math.Sqrt(100))
	fmt.Println(math.Pow(2, 3)) // pow for (2, 3) --> means 2 to the power 3
	fmt.Println(complex(2, 3))
}
