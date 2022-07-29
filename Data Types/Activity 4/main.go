package main

import "fmt"

func main(){
	var interest, principal, rate, days float64

	fmt.Println("Enter the principal:")
	fmt.Scan(&principal)
	fmt.Println("Enter the rate (as a float):")
	fmt.Scan(&rate)
	fmt.Println("Enter the days:")
	fmt.Scan(&days)
	interest = principal * rate * days / 365
	fmt.Printf("Interest: $%.2f", interest)
}