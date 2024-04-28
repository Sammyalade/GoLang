package main

import "fmt"

func main() {
	var amountEarned int
	fmt.Println("Enter amount earned for the month")
	fmt.Scanf("%d", &amountEarned)
	fmt.Printf("Tax due is %d", amountEarned)
}

func calculateTax(amountEarned int) int {
	if amountEarned <= 30000 {
		return amountEarned * (15 / 100)
	}
	return amountEarned * (20 / 100)
}
