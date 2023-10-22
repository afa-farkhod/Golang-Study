package main

import (
	"fmt"
)

var (
	balance int
)

func main() {
	fmt.Print("Enter your balance: ")
	fmt.Scanln(&balance)

	if balance < 10000 {
		fmt.Println("You can't have lunch")
	} else {
		fmt.Println("Welcome")
	}
}
