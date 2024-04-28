package main

import "fmt"

func main() {
	var number int
	var largest int
	for i := 0; i < 10; i++ {
		fmt.Print("Enter a number")
		fmt.Scanf("%d", &number)
		if number > largest {
			largest = number
		}
	}
	fmt.Printf("Highest number is %d", largest)
}
