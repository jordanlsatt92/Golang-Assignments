/*
Jordan Satterfield
6/30/22
*/

package main

import (
	"fmt"
	"os"
)

func main(){
	var desiredNumber, count int
	fmt.Print("Please enter a positive integer: ")
	fmt.Scan(&desiredNumber)
	if desiredNumber < 1{
		fmt.Println("You must enter a positive integer.")
		os.Exit(1)
	}
	count = 0
	var str string
	fmt.Println(0)
	for i:=1; count != desiredNumber; i++{
		str = ""
		if i % 3 == 0 && i % 5 == 0{
			str+= "fizz buzz"
			count++
			fmt.Println(str)
		}else if i % 3 == 0{
			str+= "fizz"
			count++
			fmt.Println(str)
		}else if i % 5 == 0{
			str+= "buzz"
			count++
			fmt.Println(str)
		}else{
			fmt.Println(i)
		}
	}
	fmt.Println("Tradition!!")
}