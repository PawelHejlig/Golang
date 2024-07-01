package main

import "fmt"

//s "strings"

func Solution(word string) string {
	runes := []rune(word)
	n := len(runes)
	fmt.Println(word, runes)
	// Reverse the slice of runes
	for i := 0; i < n/2; i++ {
		fmt.Println(i, n)
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	fmt.Println(runes)
	return string(runes)

}
