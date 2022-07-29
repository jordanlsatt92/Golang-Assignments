package main

/*
Write a program that prints the date for the next five days, starting from today.
*/
import (
	"fmt"
	"time"
)

func main(){
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02"))
	for i:=0; i < 4; i++{
		currentTime = currentTime.AddDate(0,0,1)
		fmt.Println(currentTime.Format("2006-01-02"))
	}
}