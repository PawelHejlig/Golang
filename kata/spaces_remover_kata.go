package main

import (
	//"fmt"
	s "strings"
)

func NoSpace(word string) string {
	//fmt.Println(s.ReplaceAll(word, " ", ""))
	return s.ReplaceAll(word, " ", "")
	//or
	//return s.Replace(word, " ", "", -1)
}
