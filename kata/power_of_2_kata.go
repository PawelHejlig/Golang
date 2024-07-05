package main

import (
	"fmt"
	"math"
	//s "strings"
)

func PowersOfTwo(n int) []uint64 {
	result := make([]uint64, n+1)
	for i := 0; i <= n; i++ {

		result[i] = uint64(math.Pow(2, float64(i)))
		fmt.Println(i, n, result)
	}
	return result
}

func main() {
	n := 1
	powers := PowersOfTwo(n)
	fmt.Println(powers)
}

//func PowersOfTwo(n int) (powers []uint64) {
//	for i := 0; i <= n; i++ {
//		powers = append(powers, uint64(math.Pow(2, float64(i))))
//	}
//	return
//}
