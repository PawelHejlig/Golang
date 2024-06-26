package main

import (
	"fmt"
	s "strings"
)

func AbbrevName(name string) string {
	words := s.Fields(name)
	fmt.Println(words)
	fl := make([]string, len(words))

	for i, word := range words {
		if i == len(words)-1 {
			fl[i] = s.ToUpper(string(word[0]))
		} else {
			fl[i] = s.ToUpper(string(word[0])) + "."
		}
	}
	fmt.Println(s.Join(fl, ""))
	return s.Join(fl, "")
}
