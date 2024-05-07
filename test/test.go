package main

import (
	"fmt"
)

var (
	temperature int
)

func main() {
	fmt.Print("Enter today's temperature: ")
	fmt.Scanln(&temperature)

	if temperature > 15 {
		fmt.Println("It is warm")
	} else {
		fmt.Println("It is cold")
	}
}
