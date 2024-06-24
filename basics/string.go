package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println(s.Contains("test", "e"))

	/* Build in */
	fmt.Println(len("hello")) // => 5
	// Outputs: 101 (byte value of the character 'e' is 101 in ASCII encoding)
	fmt.Println("hello"[1])
	// Outputs: e
	fmt.Println(string("hello"[0]))

	var name string
	name = "Sam Harris"
	words := s.Fields(name)
	fl := make([]string, len(words))

	for i, word := range words {
		fl[i] = string(word[0])
	}
	fmt.Println(fl)

}
