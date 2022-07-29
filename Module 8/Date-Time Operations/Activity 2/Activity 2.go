package main

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