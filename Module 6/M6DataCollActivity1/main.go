package main

// Jordan Satterfield
// Description: Create a function that takes as input an int n and returns an array of
// length n with random integers between -100 and 100.
import (
	"fmt"
	"m6DataCollActivity1/dependencies"
)

func main(){

	var input int
	fmt.Print("Please enter an integer:")
	fmt.Scan(&input)
	array:=dependencies.GenerateArray(input)
	fmt.Println("Array length =", len(array))
	fmt.Println(array)
}