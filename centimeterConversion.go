package main

// Jordan Satterfield
import (
	"fmt"
	"os"
)

func main(){

	// Length in centimeters
	var length float64 = 0
	
	// Prompt user to enter centimeter value
	fmt.Print("Please enter the length in centimeters: ")
	// Assign entered value to length variable
	fmt.Scanf("%f", &length)
	if length < 0{
		fmt.Println("You cannot enter a negative number.")
		os.Exit(1)
	}
	// Calculate the number of inches
	var inches float64 = length / 2.54

	var feet int = 0
	// Check to see if the number of inches is greater than 12
	if inches > 12 {
		// Calculate the number of feet
		feet = int(inches / 12)
	}

	// Calculate the remaining inches after dividing by length of foot
	if feet > 0{
		inches = inches - float64(feet)*12
	}

	// Check to see if the length is greater than foot 
	// Print the number of feet and inches
	if feet > 0{
		fmt.Printf("%.2f centimeters converts to %d feet and %.2f inches.", length, feet, inches)
	// Print the number of inches
	}else{
		fmt.Printf("%.2f centimeters converts to %.2f inches.", length, inches)
	}
}