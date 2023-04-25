package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// Iterate
	list := []int{2, 5, 6, 7, 8}

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	fmt.Println("---------")

	// Iterate String
	word := "iLoveU3000"
	wordCount := utf8.RuneCountInString(word)
	fmt.Println(wordCount)
	// Howwww!
	for i := 0; i < wordCount; i++ {
		fmt.Print(word[i])
	}
}
