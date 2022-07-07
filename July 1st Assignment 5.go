// Assignment 5
// Jordan Satterfield
// Description: The Fibonacci sequence is defined as : fib(0) = 0, fib(1) = 1
// fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find
// fib(n)
package main

import "fmt"

func fib(number uint) int{
	if number == 1{
		return 1
	}else if number == 0{
		return 0
	}else{
		return fib(number-1) + fib(number-2)
	}
}

func main() {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}

