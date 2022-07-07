// Assignment 3
// Jordan Satterfield
// Descsription: Write a function with one variadic parameter that finds the greatest
// number in a list of numbers.
package main

import "fmt"

func greatest(numbers... int) int{
	var largest int = 0
	for _,val:= range numbers{
		if val > largest{
			largest = val
		}
	}
	return largest
}

func main() {
	fmt.Println(greatest(10,9,21,8,7,12))
	fmt.Println(greatest(100,100,100,101,100,100))
}