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

	score := 0
	num_question := 2

	fmt.Printf("What is better, AMD or Intel? ")
	var answer string
	fmt.Scan(&answer)

	if answer == "Intel" || answer == "intel" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect, Intel is better")
	}

	fmt.Printf("How many cores does Ryzen 9 3900x have? ")
	var cores int
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect, it has 12 cores")
		//return
	}

	fmt.Printf("You scored %v out of %v.", score, num_question)
	percent := (float64(score) / float64(num_question)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
