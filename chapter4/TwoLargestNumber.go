package main

import "fmt"

func main() {
	var number int
	var largest int
	var secondLargest int

	for i := 0; i < 10; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%d", &number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest && number != largest {
			secondLargest = number
		}
	}

	fmt.Printf("Highest number is %d\n", largest)
	fmt.Printf("Second Highest number is %d", secondLargest)
}
