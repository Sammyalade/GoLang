package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Println("Enter coordinates of the first point (x1, y1):")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Enter coordinates of the second point (x2, y2):")
	fmt.Scanln(&x2, &y2)

	if x1 == x2 {
		fmt.Println("The line segment is perpendicular to the x-axis.")
	} else if y1 == y2 {
		fmt.Println("The line segment is perpendicular to the y-axis.")
	} else {
		fmt.Println("The line segment is not perpendicular to either the x-axis or the y-axis.")
	}
}
