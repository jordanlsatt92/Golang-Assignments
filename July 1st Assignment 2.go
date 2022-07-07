// Assignment 2
// Jordan Satterfield
// Description: Write a function which takes an integer and halves it and returns true if it was even or false
// if it was odd. For example, half(1) should return (0, false) and half(2) returns (1,true)
package main

import "fmt"

func half(n int) (int, bool){
	answer:= n / 2
	even:= n % 2 == 0
	return answer, even
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}