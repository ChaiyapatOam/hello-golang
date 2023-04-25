package main

import "fmt"

func main() {
	// Slice like a List
	courses := []string{}
	fmt.Println(courses)
	courses = append(courses, "MTH103", "LNG220", "CPE121", "CPE112")
	fmt.Println(courses)

	newCourse := courses[2:]
	fmt.Println(newCourse)
}
