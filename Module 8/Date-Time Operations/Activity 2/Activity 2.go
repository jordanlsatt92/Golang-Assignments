package main

/*
Write a function that prompts the user for a year and checks if it is a leap year. The function should return True if the input is a leap year and False otherwise.
*/
import (
	"fmt"
)

func isLeapYear(year int) bool{
	if year % 4 == 0{
		//Check for 
		if year % 100 == 0{
			if year % 400 == 0{
				return true
			}else{
				return false
			}
		}
		return true
	}
	return false
}

func main(){
	year1:= 2020
	year2:= 1904
	fmt.Println(year1, "leap year:", isLeapYear(year1))
	fmt.Println(year2, "leap year:", isLeapYear(year2))
}