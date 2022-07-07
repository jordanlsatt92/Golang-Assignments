// Assignment 4
// Jordan Satterfield
// Description: Using makeEvenGenerator as an example, write a makeOddGenerator function
// that generates odd numbers.
package main

import "fmt"

func makeOddGenerator() func() int{
	var number int = 1
	var temp int
	return func() int{
		temp = number
		number += 2
		return temp
	}
}

func main() {
	f:=makeOddGenerator()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}