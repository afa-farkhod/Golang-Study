package main

import (
	"fmt"
)

func printLine(length int, separator string) {
	for i := 0; i < length; i++ {
		fmt.Print(separator)
	}
	fmt.Println()
}

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]int{9, 10}
	cars := [...]string{"Nexia", "Tico", "Damas"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(cars)
	//fmt.Println(separator)
	printLine(20, "-")
	fmt.Println(cars[1])

}
