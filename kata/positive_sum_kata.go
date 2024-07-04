package main

import (
	"fmt"
	//s "strings"
)

func PositiveSum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		if numbers[i] > 0 {
			sum += numbers[i]
			fmt.Println(i, numbers[i], sum)
		}

	}
	return sum
}

//func main() {
//	PositiveSum([]int{1, -2, 3, 4, 5})
//}

//better option
//  for _, num := range numbers {
//    if num > 0 {
//		sum = sum + num
//	  }
