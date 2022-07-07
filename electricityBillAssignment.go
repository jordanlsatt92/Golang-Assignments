package main

// Jordan Satterfield
import (
	"fmt"
	"os"
)

func main(){

	// units variable to store the entered units.
	var units int
	// Prompt user for number of units
	fmt.Print("Enter the number of units: ")
	// Store entered value in units variable
	fmt.Scanf("%d", &units)

	// Check if the user has entered a valid number of units.
	if units < 0{
		fmt.Println("You cannot enter a negative number.")
		os.Exit(1)
	}

	// bill variable to store the cost of the bill
	var bill int
	// Check for values between zero and one hundred inclusive
	if units >=0 && units <=100{
		bill = units * 5
	// Check for values between 101 and 200 inclusive
	}else if units >= 101 && units <=200{
		bill = 500 + (units-100) * 7
	// Calculate bill for units greater than 200
	}else{
		bill = 1200 + (units - 200) * 10
	}

	// Print the bill to the console
	fmt.Printf("\nYour bill is $%d dollars.", bill)
}