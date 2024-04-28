package main

import "fmt"

func main() {
	fmt.Println("N\tN2\tN3\tN4")
	for i := 1; i < 6; i++ {
		fmt.Printf("%d\t%d\t%d\t%d\n", i, i*i, i*i*i, i*i*i*i)
	}
}
