package main

import (
	//s "strings"
	"fmt"
	"unicode"
)

func ToAlternatingCase(str string) string {
	result := []rune(str)
	for i, r := range result {

		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			result[i] = unicode.ToUpper(r)
		}
		fmt.Println(i, r, result[i])
	}
	fmt.Println(result)
	return string(result)
}

func main() {
	ToAlternatingCase("HEllo world")
}
