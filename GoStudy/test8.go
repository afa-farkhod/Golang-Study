package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"Go", "Slices", "Are", "Popular"}
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)
}
