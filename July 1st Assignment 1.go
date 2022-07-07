// Jordan Satterfield
// Assignment 1
// Description: Sum is a function which takes a slice of numbers and
// adds them together. What would its function signature look like in Go?

package main

import "fmt"

func sum(numbers []int) int {
	var total int
	for _, val := range numbers {
		total += val
	}
	return total
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(sum(slice))
	slice2 := []int{1, 2, -1, -2}
	fmt.Println(sum(slice2))
}
