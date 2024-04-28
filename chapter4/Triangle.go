package main

import (
	"fmt"
)

func main() {
	var number int
	print("Enter number of triangle line")
	fmt.Scanf("%d", &number)
	for asterisks := 1; asterisks <= number; asterisks++ {
		for asterisks2 := 1; asterisks2 <= asterisks; asterisks2++ {
			print("*")
			print(" ")

		}
		println()
	}

}
