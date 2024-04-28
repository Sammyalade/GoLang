package main

import (
	"fmt"
)

func main() {
	var value int
	fmt.Print("Enter a value number: ")
	fmt.Scanln(&value)

	sum := 0

	for sum < value {
		var num int
		fmt.Print("Enter an integer: ")
		fmt.Scanln(&num)

		sum += num
		fmt.Printf("Current sum: %d\n", sum)
	}

	fmt.Println("Sum has reached or exceeded the value number!")
}
