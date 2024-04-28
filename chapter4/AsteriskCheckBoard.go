package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter board size: ")
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
