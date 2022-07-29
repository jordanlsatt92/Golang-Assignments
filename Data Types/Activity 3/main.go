package main

import (
	"fmt"
	"math"
)

func main(){
	var P, r, n, t float64
	fmt.Print("Please enter inital deposit:")
	fmt.Scan(&P)
	fmt.Print("Please enter interest rate (as a float):")
	fmt.Scan(&r)
	fmt.Print("Please enter the number of times per year interest is calculated:")
	fmt.Scan(&n)
	fmt.Print("Please enter the number of years since initial deposit:")
	fmt.Scan(&t)

	fmt.Println("Initial deposit:", P)
	fmt.Println("Interest rate:", r)
	fmt.Println("The number of times per year interest is calculated:", int(n))
	fmt.Println("Number of years since initial deposit:", int(t))
	var value float64 = P * math.Pow((1 + (r/n)), n*t)
	fmt.Println("Value:", value)
}