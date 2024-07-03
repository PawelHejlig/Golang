package main

import (
	"fmt"
	//s "strings"
)

func Summation(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		fmt.Println(i, sum)
	}
	return sum
}

//func main() {
//	Summation(8)
//}

//func Summation(n int) int {
//    return n * (n + 1) / 2
//}
