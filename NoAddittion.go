package main

import "fmt"

func main() {
	fmt.Print(noAddition(6, -9))
}
func noAddition(num int, num2 int) int {
	for i := 0; i < num; i++ {
		num2++
	}
	return num2
}
