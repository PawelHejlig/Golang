package main

import "fmt"

/*import "sort"


func ExpressionMatter(a int, b int, c int) int {
    arr := []int{a*(b+c), a*b*c, a+b+c, (a+b)*c}
    sort.Ints(arr)
    return arr[len(arr)-1]
}
*/

func ExpressionMatter(a int, b int, c int) int {
	x1 := (a * (b + c))
	x2 := (a * b * c)
	x3 := (a + b*c)
	x4 := ((a + b) * c)
	x5 := (a + b + c)

	numbers := []int{x1, x2, x3, x4, x5}
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
	return max
}
