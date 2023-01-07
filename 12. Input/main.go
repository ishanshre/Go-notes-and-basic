package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create a reusuable input function with bufio
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create  a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Bill created - ", b.name)
	return b
}

func prmptOptions(b bill) {
	reader := bufio.NewReader((os.Stdin))
	opt, _ := getInput("choose options (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		fmt.Println("yopur choice a")
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Item Price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("This must be number")
			prmptOptions(b)
		}
		b.addItem(name, p)
		fmt.Printf("Item added -: %v --> $%v", name, p)
		prmptOptions(b)
	case "t":
		tip, _ := getInput("Entert the tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("This must be number")
			prmptOptions(b)
		}
		b.update(t)
	case "s":
		b.save()
		fmt.Println("You choose S")
	default:
		fmt.Println("not a valid choice")
		prmptOptions(b)
	}
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	prmptOptions(mybill)
	fmt.Println(mybill)
}
