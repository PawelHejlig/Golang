package main

import "fmt"

//s "strings"

func countSheep(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		fmt.Println(i, num)
		result += fmt.Sprintf("%d sheep...", i)
	}

	return result
}

/*func main() {
	countSheep(2)
}

/*second option
func countSheep(num int) string {
	var sb strings.Builder

	for count := 1; count <= num; count++ {
		fmt.Fprintf(&sb, "%d sheep...", count)
	}

	return sb.String()
}
*/
