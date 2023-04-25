package main

import "fmt"

func zero(val int) {
	val = 0
}

func zeroPointer(valPointer *int) {
	*valPointer = 0
}

func main() {
	i := 1

	zero(i)
	fmt.Println("i =", i)

	zeroPointer(&i)
	fmt.Println("iPointer = ", i)
}
