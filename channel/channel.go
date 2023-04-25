package main

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, ":", i)
	}
}
func main() {
	go f("Hello")
	go f("CPE112121")
	time.Sleep(5 * time.Second)
}
