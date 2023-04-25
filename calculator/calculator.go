package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Print("", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must be a number", prompt)
		panic(message)
	}
	return value
}
func getOperator() string {
	fmt.Println("+ - * /")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(x, y float64) float64 {
	return x + y
}

func main() {
	val1 := getInput("value =")
	val2 := getInput("value =")

	var result float64
	operator := getOperator()
	switch operator {
	case "+":
		result = add(val1, val2)
	case "-":
		result = add(val1, val2)
	default:
		panic("ERROR")
	}

	fmt.Println("result =", result)

}
