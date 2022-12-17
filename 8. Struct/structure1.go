package main

import "fmt"

type Student struct {
	name  string
	age   int
	class int
}

func main() {
	var stu1 Student
	stu1.name = "Roshan"
	stu1.age = 10
	stu1.class = 7
	var stu2 Student
	stu2.name = "Bijay"
	stu2.age = 12
	stu2.class = 9
	fmt.Printf("Name ---> %v\n", stu1.name)
	fmt.Printf("Age ---> %v\n", stu1.age)
	fmt.Printf("Class ---> %v\n", stu1.class)
	fmt.Printf("Name ---> %v\n", stu2.name)
	fmt.Printf("Age ---> %v\n", stu2.age)
	fmt.Printf("Class ---> %v\n", stu2.class)
	// passing struct to the function
	var stu3 Student
	stu3.name = "Nishcal"
	stu3.age = 9
	stu3.class = 2
	passFunc(stu3)
	passFunc(stu2)
}

func passFunc(stu Student) {
	fmt.Printf("Name ---> %v\n", stu.name)
	fmt.Printf("Age ---> %v\n", stu.age)
	fmt.Printf("Class ---> %v\n", stu.class)
}
