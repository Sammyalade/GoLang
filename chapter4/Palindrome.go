package main

import (
	"fmt"
	"strconv"
)

func isFiveDigitPalindrome(num int) bool {
	if num < 10000 || num > 99999 {
		return false
	}

	numStr := strconv.Itoa(num)

	if numStr[0] == numStr[4] && numStr[1] == numStr[3] {
		return true
	}

	return false
}

func main() {
	var num int
	fmt.Println("Enter five digits: ")
	fmt.Scanf("%d", &num)
	isPalindrome := isFiveDigitPalindrome(num)
	if isPalindrome {
		fmt.Printf("%d is a palindrome", num)
	} else {
		fmt.Printf("%d is not a palindrome", num)
	}
}
