package main

import "fmt"

type Person struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	isAdmin bool
}

func main() {
	// Type Struct
	p := Person{
		Name:    "Kabigon",
		Age:     "23",
		isAdmin: false,
	}
	fmt.Println(p, p.Name)
}