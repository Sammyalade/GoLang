package main

import "fmt"

func main() {
	var goodsSold int
	fmt.Println("Enter amount of goods sold for the month")
	fmt.Scanf("%d", &goodsSold)
	fmt.Printf("Salary for the month is: %d\n", calculateSalary(goodsSold))
}

func calculateSalary(goodsSold int) int {
	commission := float64(goodsSold) * 0.09
	totalSalary := int(commission) + 200
	return totalSalary
}
