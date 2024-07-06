package main

//import "fmt"

//"fmt"
//s "strings"

func CountBy(x, n int) []int {
	var numbers []int
	//numbers[0] = x
	for i := 1; i <= n; i++ {
		//fmt.Println(i, x, n)
		numbers = append(numbers, i*x)
	}
	//fmt.Println(numbers)
	return numbers
}

/*func main() {
	CountBy(2, 5)
}

//better option
func CountBy(x, n int) (a []int) {
	for i := 1; i <= n; i++ {
		a = append(a, x*i)
	}
	return a
}
*/
