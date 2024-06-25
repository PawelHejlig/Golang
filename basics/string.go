package main

import (
	"fmt"
	s "strings"
)

func main() {
	var name string
	name = "Sam Harris"
	words := s.Fields(name)
	fl := make([]string, len(words))

	for i, word := range words {
		if i == len(words)-1 {
			fl[i] = string(word[0])
		} else {
			fl[i] = string(word[0]) + "."
		}
		fmt.Println(fl)

		fmt.Println(s.Join(fl, " "))
	}
}
