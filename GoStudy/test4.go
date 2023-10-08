package main

import "fmt"

var (
	num1 int
	num2 int
)

func addition(a int, b int) int {
	return a + b
}

func main() {
	fmt.Print("Enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter number 2: ")
	fmt.Scanln(&num2)
	sum := addition(num1, num2)
	fmt.Println("Sum=", sum)
}
