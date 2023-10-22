package main

import (
	"fmt"
)

var (
	name string
	age  int
)

func main() {
	fmt.Print("Hi, welcome:) What is your name?: ")
	fmt.Scanln(&name)

	fmt.Println("Welcome ", name)

	fmt.Println("Now tell my your age: ")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("You are too young to have sex")
	} else {
		fmt.Println("Enjoy the party")
	}
}
