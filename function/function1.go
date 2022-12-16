package main

import "fmt"

func main() {
	fullName("ishan", "shrestha")
}

func fullName(fname string, lname string) {
	fmt.Println(fname + " " + lname)
}
