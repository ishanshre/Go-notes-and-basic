package main

import "fmt"

func main() {
	mybill := newBill("marios", map[string]float64{"pizza": 10.2, "coffe": 5.8})
	mybill.update(10.1)
	mybill.addItem("milk", 2.2)
	fmt.Println(mybill.format())
}
