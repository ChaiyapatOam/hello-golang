package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func plus(x int, y int) int {
	result := x + y
	return result
}
func sum(x, y, z int) int {
	return x + y + z
}

func main() {
	hello()
	x := plus(3, 5)
	fmt.Println(x)
}
