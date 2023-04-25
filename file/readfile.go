package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("New-Year-Resolution.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
}
