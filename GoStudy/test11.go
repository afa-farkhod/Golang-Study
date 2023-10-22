package main

//import "fmt"

import (
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//original slice
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	//copy slice
	neededNumbers := numbers[:len(numbers)-3]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	//print the result
	fmt.Println(numbersCopy)
	fmt.Println(len(numbersCopy))
	fmt.Println(cap(numbersCopy))
}
