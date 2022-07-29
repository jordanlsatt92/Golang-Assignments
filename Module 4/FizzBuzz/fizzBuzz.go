/*
Jordan Satterfield
6/30/22

    Ask the user for a number.
    Output a count starting with 0.
        Display the count number if it is not divisible by 3 or 5.
        Replace every multiple of 3 with the word "fizz."
        Replace every multiple of 5 with the word "buzz."
        Replace multiples of both 3 and 5 with "fizz buzz."
    Continue counting until the number of integers replaced with "fizz," "buzz," or "fizz buzz" reaches the input number.
    The last output line should read "TRADITION!!"

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