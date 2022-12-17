package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(getName())
	fmt.Printf("multiple =\t %v\n", multiple(2, 4))
	result := multiple(5, 6)
	fmt.Println(result)
	//directly printing more than one returned variable
	fmt.Println(div1(6, 4)) // output 1 2  ---> 1 - resukt and 2 ----> remain
	// assigning reutrned values to the  variable
	result, remain := div1(5, 2)
	fmt.Printf("result => %v \t remain => %v\n", result, remain)
	// We can also omit the returned values using underscore
	result1, _ := div1(5, 2)
	fmt.Printf("result --> %v\n", result1)
	_, remain2 := div1(5, 2)
	fmt.Printf("remain --> %v\n", remain2)
}

func sum(x int, y int) int {
	return x + y
}

// for return finction we must define what type of datatype is returned

func getName() string {
	return "IShan Srestha"
}

// named returned function. We use return statemement without specifing variable

func multiple(x int, y int) (result int) {
	result = x * y
	return
}

// returning multiple results

func div1(x int, y int) (result int, remain int) {
	result = x / y
	remain = x % y
	return
}
