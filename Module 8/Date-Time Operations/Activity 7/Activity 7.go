package main

/*
Write a program that adds five seconds to the current time and displays the result.
*/
import (
	"fmt"
	"time"
)

func main(){
	currentTime:= time.Now()
	newTime := currentTime.Add(5*time.Second)
	fmt.Println("Current time:",currentTime)
	fmt.Println("5 seconds later:", newTime)
	
}