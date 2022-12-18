package main

import "fmt"

func handlePac() {
	a := recover()
	if a != nil {
		fmt.Println("Recover", a)
	}
}

func main() {
	defer handlePac()
	var num1 float32
	var num2 float32
	var result float32
	fmt.Println("Enter two numbers: ")
	fmt.Scan(&num1, &num2)
	if num2 == 0 {
		panic("Second number input cannot be zero")
	}
	result = num1 / num2
	fmt.Printf("the result of %v divided by %v is %v \n", num1, num2, result)
}
