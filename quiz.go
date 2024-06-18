package main

import "fmt"

func main() {

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Jej you can play")
	} else {
		fmt.Println("You can't play")
		return
	}

	fmt.Printf("What is better, AMD or Intel? ")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)
}
