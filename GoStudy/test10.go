package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
