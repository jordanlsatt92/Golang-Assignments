package main

/*
Write a program to display the following:

    Current date and time
    Current year
    Current month
    Week number of the year
    Weekday of the week
    Day of the year
    Day of the month
    Day of the week
*/
import (
	"fmt"
	"time"
)

func main(){
	currentTime:= time.Now()
	fmt.Println("Current date and time:", currentTime.Format("01-02-2006 04:03:02"))
	fmt.Println("Current Year:", currentTime.Year())
	fmt.Println("Current Month:", currentTime.Month())
	_,week:= currentTime.ISOWeek()
	fmt.Println("Week number of year:", week )
	fmt.Println("Weekday:", currentTime.Weekday())
	fmt.Println("Day of the year:", currentTime.YearDay())
	fmt.Println("Day of the month:", currentTime.Day())
	var day int 
	switch(currentTime.Weekday()){
	case time.Sunday:
		day = 1
	case time.Monday:
		day = 2
	case time.Tuesday:
		day = 3
	case time.Wednesday:
		day = 4
	case time.Thursday:
		day = 5
	case time.Friday:
		day = 6
	case time.Saturday:
		day = 7
	}

	fmt.Println("Day of the week:", day)
}