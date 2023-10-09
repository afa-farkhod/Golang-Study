package main

import (
	"fmt"
)

var (
	number1 int
	number2 int
	number3 int
)

func sum(a int, b int, c int) int {
	return a + b + c
}

func main() {
	fmt.Print("Enter number 1: ")
	fmt.Scanln(&number1)
	fmt.Print("Enter number 2: ")
	fmt.Scanln(&number2)
	fmt.Print("Enter number 3: ")
	fmt.Scanln(&number3)

	result := sum(number1, number2, number3)

	fmt.Println("Sum = ", result)
}
