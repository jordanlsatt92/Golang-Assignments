package main

/*
Write a program that prompts the user for two different dates, computes the number of days between those dates, and displays the results.

The order in which the user enters the dates should not affect the results.
*/
import (
	"fmt"
	"math"
	"time"
)

func main(){
	var year1,day1 int
	var month1 time.Month
	fmt.Print("Please enter a year:")
	fmt.Scan(&year1)
	fmt.Print("Please enter a month (1-12):")
	fmt.Scan(&month1)
	fmt.Print("Please enter a day (1-31):")
	fmt.Scan(&day1)
	inputTime1 := time.Date(year1,month1,day1,0,0,0,0, time.Now().Location())

	var year2,day2 int
	var month2 time.Month
	fmt.Print("Please enter a year:")
	fmt.Scan(&year2)
	fmt.Print("Please enter a month (1-12):")
	fmt.Scan(&month2)
	fmt.Print("Please enter a day (1-31):")
	fmt.Scan(&day2)
	inputTime2 := time.Date(year2,month2,day2,0,0,0,0, time.Now().Location())

	daysBetween := int(math.Abs(float64(inputTime2.Sub(inputTime1).Hours() / 24)))

	fmt.Println(daysBetween, "days in between")
	
}