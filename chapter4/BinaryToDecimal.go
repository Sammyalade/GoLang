package main

import (
	"fmt"
)

func binaryToDecimal(binary string) int {
	decimal := 0
	power := 1

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += power
		}
		power *= 2
	}

	return decimal
}

func main() {
	var binary string
	fmt.Print("Enter a binary number (containing only 0s and 1s): ")
	fmt.Scanln(&binary)

	for char := 0; char < len(binary); char++ {
		if binary[char] != '0' && binary[char] != '1' {
			fmt.Println("Invalid binary number.")
			return
		}
	}
	fmt.Printf("Decimal equivalent: %d\n", binaryToDecimal(binary))
}
