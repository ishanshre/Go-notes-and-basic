package main

import "fmt"

func main() {
	a := make(map[int]string)
	fmt.Println(a)
	fmt.Println("Size of the map a is ", len(a))
	a[1] = "Binod"
	a[2] = "Tamang"
	fmt.Println(a)
	fmt.Println("Size of map a after assigining valueis :", len(a))
	for key, value := range a {
		fmt.Printf("%d ---> %v\n", key, value)
	}
	b := make(map[string]int)
	fmt.Println("map b declare: ", b)
	b["name"] = 1
	fmt.Println("map b is : ", b["name"])
	for k, v := range b {
		fmt.Println(k, v)

	}
	for _, v := range b {
		fmt.Println(v)
	}
	for k, _ := range b {
		fmt.Println(k)
	}
	c := map[int]string{1: "a", 2: "b"}
	fmt.Println(c)
	for k, v := range c {
		fmt.Println(k, " ", v)
	}
}
