package main

import (
	"fmt"
	s "strings"
)

func StringToArray(str string) []string {
	fmt.Println(str)
	words := s.Fields(str)
	fmt.Println(words)
	return words
}

//func main() {
//	StringToArray("John Dou")
//}

// return strings.Fields(str)
