package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

// json.Marshal key should be start with capital letter
func main() {
	data, _ := json.Marshal(&employee{20, "asdfdsfsd", "dsfdsf", "dsfdsf"})
	fmt.Println(string(data))
	// Unmarshal
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":10,"Name":"Kabigon", "Tel" : "01235", "Email":"Chaiyapat@7707"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
}
