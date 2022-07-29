package main

import "fmt"

func main(){
	var born, live string
	fmt.Print("Enter the state you were born: ")
	fmt.Scan(&born)
	fmt.Print("Enter the state where you live: ")
	fmt.Scan(&live)
	fmt.Println("Born State:", born)
	fmt.Println("Current State:", live)
}