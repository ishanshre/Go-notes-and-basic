package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0.0 {
		return result, errors.New("cannot divide by zerro")
	}
	result = x / y
	return result, nil
}
