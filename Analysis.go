package main

import "fmt"

func main() {
	var passes int
	var failures int

	for i := 0; i < 10; i++ {
		var result int
		fmt.Print("Enter result (1 = pass, 2 = fail): ")
		fmt.Scanf("%d", result)

		if result == 1 {
			passes += 1
		} else if result == 2 {
			failures += 1
		} else {
			fmt.Print("Incorrect result")
			i--
		}
	}
	if passes >= 8 {
		fmt.Print("Bonus to instructor")
	}
}
