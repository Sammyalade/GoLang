package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var input int
	fmt.Scanln(&input)
	output := 1

	for i := input; i > 0; i-- {
		output *= i
	}
	fmt.Println(output)
}
