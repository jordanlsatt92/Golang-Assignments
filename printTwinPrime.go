package main

import "fmt"

func main(){

	var isPrime bool = true
	var primeTwin bool = true

	for i:=2; i <=100; i++{
		isPrime = true
		primeTwin = true
		for j:=2; j < i; j++{
			if i % j == 0{
				isPrime = false
				break;
			}
		}

		for j:=2; j < i+2; j++{
			if ((i+2) % j == 0){
				primeTwin = false
				break
			}
		}
		if isPrime && primeTwin{
			fmt.Println(i, "and", (i+2), "are twin primes.")
		}

	}
}