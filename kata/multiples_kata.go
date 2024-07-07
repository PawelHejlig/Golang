package main

import "fmt"

func FindMultiples(integer, limit int) (result []int) {
	if limit%integer == 0 && integer != 0 && limit != 0 {
		for i := integer; i <= limit; i += integer {
			result = append(result, i)
			fmt.Println(integer, limit, limit%integer, result)
		}
		return result
	}
	return nil
}

func main() {
	FindMultiples(5, 7)
}

//better option
//  func FindMultiples(integer, limit int) (res[]int) {
//  for i := integer; i <= limit; i += integer { res = append(res, i)}
//  return
//}
