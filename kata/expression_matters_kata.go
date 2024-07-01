package main

import "fmt"

//import "fmt"

//s "strings"

func ExpressionMatter(a int, b int, c int) int {
	var x1, x2, x3, x4 int
	x1 = (a * (b + c))
	x2 = (a * b * c)
	x3 = (a + b*c)
	x4 = ((a + b) * c)

	numbers := []int{x1, x2, x3, x4}
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
	return max
}

func main() {
	ExpressionMatter(1, 2, 3)
}
