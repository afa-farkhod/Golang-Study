package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter your name: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Now enter your surname: ")
	var surname string
	fmt.Scanln(&surname)

	fullName := username + " " + surname
	fmt.Println("Your full name is: ", fullName)
}
