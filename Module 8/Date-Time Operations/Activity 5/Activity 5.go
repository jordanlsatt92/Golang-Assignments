package main

/*
Write a program that prints the dates for yesterday, today, and tomorrow.
*/
import (
	"fmt"
	"time"
)

func main(){
	currentTime := time.Now()
	yesterday := currentTime.AddDate(0,0,-1)
	tomorrow := currentTime.AddDate(0,0,1)

	fmt.Println("Yesterday:", yesterday.Format("2006-01-02"))
	fmt.Println("Today:", currentTime.Format("2006-01-02"))
	fmt.Println("Tomorrow:", tomorrow.Format("2006-01-02"))
}