package main

import (
	"fmt"
)

var (
	xValue     int
	yValue     int
	xyAddition int
)

func main() {
	fmt.Print("x=")
	fmt.Scanln(&xValue)

	fmt.Print("y=")
	fmt.Scanln(&yValue)

	xyAddition = xValue + yValue

	fmt.Println("x + y =", xyAddition)
}
