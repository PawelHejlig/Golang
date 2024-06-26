package main

import (
	"fmt"
	s "strings"
)

func NoSpace(word string) []string {
	words := s.Fields(word)
	fmt.Println(words)
	return words
}
