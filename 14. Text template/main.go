package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func main() {
	fmt.Println("Hello World")
	secret := secret{
		Name:   "Ishan",
		Secret: "Shrestha",
	}
	templatePath := "/home/ishan/project/Go-notes-and-basic/14. Text template/template.txt"
	t, err := template.New("template.txt").ParseFiles(templatePath)
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Println(err)
	}
}
