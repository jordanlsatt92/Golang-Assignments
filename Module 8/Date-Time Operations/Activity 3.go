package main

import (
	"fmt"
	"time"
)

func main(){

	var year,day int
	var month time.Month
	fmt.Print("Please enter a year:")
	fmt.Scan(&year)
	fmt.Print("Please enter a month (1-12):")
	fmt.Scan(&month)
	fmt.Print("Please enter a day (1-31):")
	fmt.Scan(&day)
	inputTime := time.Date(year,month,day,0,0,0,0, time.Now().Location())
	changedTime:= inputTime.AddDate(0,0,-5)
	

	fmt.Println("Original date entered:", inputTime)
	fmt.Println("Original - 5 days:", changedTime)
}
