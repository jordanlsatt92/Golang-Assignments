package main

import (
	"fmt"
	"regexp"
)
func checkString(str string) bool{
	match,err := regexp.MatchString("^[a-zA-Z0-9]+", str)
	if err != nil{
		fmt.Println("Invalid input. Please try again")
		return false
	}
	return match
}

func main(){
	str1:="AskdjDG90SD"
	str2:="!asD90>>"
	str3:= ""
	fmt.Println(str1, " only contains numbers and letters:", checkString(str1))
	fmt.Println(str2, " only contains numbers and letters:", checkString(str2))
	fmt.Println(str3, " only contains numbers and letters:", checkString(str3))
}