package main

import (
	"encoding/json"
	"fmt"
)

type Warrior struct {
	Name   string
	Weapon string
	Level  int
}

func main() {
	jsonStr := `{"Name":"Ishan","Weapon":"sword","Level":99}`
	fmt.Println(jsonStr)
	fmt.Printf("%T : %v \n", jsonStr, jsonStr)
	var shrestha Warrior
	err := json.Unmarshal([]byte(jsonStr), &shrestha)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(shrestha)
	fmt.Printf("%T : %v \n", shrestha, shrestha)

	binod := Warrior{
		Name:   "Binod",
		Weapon: "Knife",
		Level:  99,
	}
	convertToJson, err := json.Marshal(binod)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(convertToJson)
	fmt.Printf("%s:\n ", convertToJson)
	fmt.Printf("%T : \n", convertToJson)
}
