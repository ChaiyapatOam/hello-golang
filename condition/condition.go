package main

import "fmt"

func Grading(score int) string {
	var grade string
	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "B"
	} else if score >= 50 {
		grade = "C"
	} else {
		grade = "F"
	}
	return grade
}

func main() {
	var score int
	fmt.Println("Grade Calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("Score = ", score)
	fmt.Println("Grade = ", Grading(score))
}
