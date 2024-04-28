package main

import "fmt"

func main() {
	var accountNumber int
	var openingBalance, totalItemCharged, creditAmount, creditLimit int

	fmt.Println("Enter Account Number:")
	fmt.Scanf("%d", &accountNumber)

	fmt.Println("Enter Opening Balance:")
	fmt.Scanf("%d", &openingBalance)

	fmt.Println("Enter Total Items Charged:")
	fmt.Scanf("%d", &totalItemCharged)

	fmt.Println("Enter Total Credit Applied to Customer Account:")
	fmt.Scanf("%d", &creditAmount)

	fmt.Println("Enter Credit Limit to Customer Account:")
	fmt.Scanf("%d", &creditLimit)

	newBalance := calculateBalance(openingBalance, totalItemCharged, creditAmount)
	result := display(creditLimit, newBalance)

	fmt.Println(result)
}

func calculateBalance(openingBalance, charges, credits int) int {
	return openingBalance + charges - credits
}

func display(limit, newBalance int) string {
	if newBalance <= limit {
		return "Limit not exceeded"
	}
	return "Credit limit exceeded"
}
