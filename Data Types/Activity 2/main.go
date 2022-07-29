package main

import "fmt"

func main() {
	var input float64
	fmt.Println("Enter a float number:")
	fmt.Scan(&input)
	fmt.Println("Int portion:",int(input))

}