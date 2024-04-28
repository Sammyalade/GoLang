package main

import "fmt"

func main() {

	var milesDriven float32
	var gallonUsed float32
	var totalMilesPerGalon float32
	for milesDriven != -1 {
		fmt.Print("Enter Miles Driven or -1 to cancel: ")
		fmt.Scanf("%f", &milesDriven)
		if milesDriven == -1 {
			break
		}
		fmt.Println("Enter the Galon Used: ")
		fmt.Scanf("%f", &gallonUsed)
		totalMilesPerGalon += milesDriven / gallonUsed
	}
	fmt.Printf("Total Miles Per Gallon equals %f", totalMilesPerGalon)
}
