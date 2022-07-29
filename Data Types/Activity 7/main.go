package main

import "fmt"

func main(){
	var num1, num2, num3, num4, num5 int

	fmt.Print("Enter an integer:")
	fmt.Scan(&num1)
	fmt.Print("Enter an integer:")
	fmt.Scan(&num2)
	fmt.Print("Enter an integer:")
	fmt.Scan(&num3)
	fmt.Print("Enter an integer:")
	fmt.Scan(&num4)
	fmt.Print("Enter an integer:")
	fmt.Scan(&num5)

	fmt.Println("Number entered:", num1, num2, num3, num4, num5)
	fmt.Println("Product of numbers:", num1 * num2 * num3 * num4 * num5)
	fmt.Println("Sum of numbers:", num1 + num2 + num3 + num4 + num5)
	fmt.Println("Average of numbers:", float64((num1 + num2 + num3 + num4 + num5) / 5))

}