package main

import (
	"fmt"
	"math"
	"strconv"
)

func main(){
	var number int64
	fmt.Print("Enter an integer:")
	fmt.Scan(&number)

	fmt.Println("You selected:", number)
	val,_ := strconv.ParseBool(strconv.Itoa(int(number)))
	fmt.Println("Boolean:", val)
	fmt.Println("Binary:", strconv.FormatInt(number, 2))
	fmt.Println("Square root:", math.Sqrt(float64(number)))
}